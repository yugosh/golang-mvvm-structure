package controllers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"BACKEND-GOLANG-MVVM/internal/app/models"
	"BACKEND-GOLANG-MVVM/internal/app/services"
	"BACKEND-GOLANG-MVVM/internal/app/viewmodels"
	"BACKEND-GOLANG-MVVM/internal/config"

	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	UserService services.UserService
}

// GoogleLogin mengarahkan pengguna ke halaman login Google OAuth
func (a *AuthController) GoogleLogin(c *gin.Context) {
	state := "login"
	url := viewmodels.GoogleOauthConfig.AuthCodeURL(state)

	fmt.Println("Redirecting to Google OAuth (login):", url)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleRegister mengarahkan pengguna ke halaman registrasi Google OAuth
func (a *AuthController) GoogleRegister(c *gin.Context) {
	state := "register"
	url := viewmodels.GoogleOauthConfig.AuthCodeURL(state)

	fmt.Println("Redirecting to Google OAuth (register):", url)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback menangani callback dari Google OAuth
func (a *AuthController) GoogleCallback(c *gin.Context) {
	receivedState := c.Query("state")

	if receivedState != "login" && receivedState != "register" {
		fmt.Println("Invalid OAuth state")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	code := c.Query("code")
	token, err := viewmodels.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Code exchange failed:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	client := viewmodels.GoogleOauthConfig.Client(context.Background(), token)
	userInfoResp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer userInfoResp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(userInfoResp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
		return
	}

	googleID := userInfo["id"].(string)
	email := userInfo["email"].(string)

	if receivedState == "login" {
		// Logika untuk login
		existingUser, err := a.UserService.GetUserByGoogleID(googleID)
		if err != nil || existingUser == nil {
			// existingUser, _ = a.UserService.GetUserByEmail(email)

			existingUser, err = a.UserService.GetUserByEmail(email)
			if err != nil {
				// Tangani error di sini
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user by email"})
				return
			}
		}

		if existingUser != nil {
			userIDStr := strconv.FormatUint(uint64(existingUser.ID), 10)
			err = config.RedisClient.Set(context.Background(), userIDStr, "logged_in", 0).Err()
			if err != nil {
				log.Printf("Failed to store google ID in Redis: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store session in Redis"})
				return
			}

			// Redirect ke dashboard jika login berhasil
			c.Redirect(http.StatusFound, "/dashboard?status=success&id="+userIDStr)
			// c.JSON(http.StatusOK, gin.H{
			// 	"message": "User logged in successfully",
			// 	"user":    existingUser,
			// })
			return
		}

		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not registered. Please register first.",
		})
		return
	}

	if receivedState == "register" {
		// Logika untuk registrasi
		existingUser, err := a.UserService.GetUserByGoogleID(googleID)
		if err == nil && existingUser != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": "User already registered. Please log in.",
			})
			return
		}

		username := googleID
		// password := generateRandomHash()
		password := "12345"

		user := &models.User{
			Username:  username,
			Email:     email,
			GoogleID:  googleID,
			Password:  password,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := a.UserService.Register(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User registered successfully",
			"user":    user,
		})
		return
	}
}

// generateRandomHash menghasilkan password hash random
func generateRandomHash() string {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		panic(err)
	}
	password := base64.URLEncoding.EncodeToString(randomBytes)

	// Hash the random password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
