// Di file internal/app/controllers/user_controller.go, kita akan membuat handler untuk register dan login.

package controllers

import (
	"BACKEND-GOLANG-MVVM/internal/app/mappers"
	"BACKEND-GOLANG-MVVM/internal/app/models"
	"BACKEND-GOLANG-MVVM/internal/app/services"
	"BACKEND-GOLANG-MVVM/internal/config"
	"context"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (c *UserController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UserService.Register(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, mappers.ToUserViewModel(&user))
}

func (c *UserController) Login(ctx *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.UserService.Login(loginData.Email, loginData.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userIDStr := strconv.FormatUint(uint64(user.ID), 10)
	err = config.RedisClient.Set(context.Background(), userIDStr, "logged_in", 0).Err()
	if err != nil {
		log.Printf("Failed to store user ID in Redis: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store session"})
		return
	}

	ctx.JSON(http.StatusOK, mappers.ToUserViewModel(user))
}

func (c *UserController) Logout(ctx *gin.Context) {
	// Misalnya, kita mengambil ID pengguna dari session atau token yang sedang aktif
	userIDStr := ctx.Param("userID") // atau Anda bisa mengambilnya dari query, body, atau session

	// Hapus user ID dari Redis
	err := config.RedisClient.Del(context.Background(), userIDStr).Err()
	if err != nil {
		log.Printf("Failed to delete user ID from Redis: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove session"})
		return
	}

	// Respon berhasil logout
	ctx.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}

func (c *UserController) SendToQueue(ctx *gin.Context) {
	var requestData struct {
		Message  string `json:"message" binding:"required"`
		SocketID string `json:"socketId" binding:"required"` // Tambahkan socketId ke struct
	}

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	delay := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(delay)

	err := services.PublishMessage("test_queue", requestData.SocketID, requestData.Message)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message to queue"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Message sent to queue successfully",
		"socketId": requestData.SocketID,
	})
}

func (c *UserController) GetLastMessage(ctx *gin.Context) {
	lastMessage := services.GetLastReceivedMessage()
	ctx.JSON(http.StatusOK, gin.H{"message": lastMessage})
}
