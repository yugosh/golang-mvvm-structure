// Di file internal/app/services/user_service.go, kita akan membuat layanan untuk register dan login.

package services

import (
	"BACKEND-GOLANG-MVVM/internal/app/models"
	"BACKEND-GOLANG-MVVM/internal/app/repositories"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func init() {
	godotenv.Load()
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) Register(user *models.User) error {
	// Hash the password
	if err := user.HashPassword(); err != nil {
		return err
	}

	// Check if email or username is already taken
	if _, err := s.UserRepo.FindByEmail(user.Email); err == nil {
		return errors.New("email already registered")
	}
	if _, err := s.UserRepo.FindByUsername(user.Username); err == nil {
		return errors.New("username already taken")
	}

	// Save the user
	return s.UserRepo.CreateUser(user)
}

func (s *UserService) Login(email, password string) (*models.User, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Debugging log: Tampilkan password yang di-hash di database
	log.Printf("Hashed password from DB: %s", user.Password)

	// Debugging log: Tampilkan password yang diinput oleh pengguna
	log.Printf("Password provided: %s", password)

	// Compare the hashed password in the database with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Println("Password mismatch") // Debugging log: Password tidak cocok
		return nil, errors.New("incorrect password")
	}

	log.Println("Password match") // Debugging log: Password cocok

	files := []string{"attachment/test.pdf", "attachment/images01.jpg"}

	// Call the SendLoginNotification function after successful login
	if err := s.SendLoginNotification(email, files); err != nil {
		log.Printf("Failed to send login notification: %v", err)
		// Optional: decide whether to return the error or just log it
	}

	return user, nil
}

func (vm *UserService) SendLoginNotification(email string, filePaths []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USERNAME"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Login Berhasil")

	// Set body email dengan konten HTML
	htmlContent := `
        <html>
        <body>
            <h1>Login Berhasil</h1>
            <p>Anda berhasil login ke sistem kami. Silakan periksa lampiran untuk informasi lebih lanjut.</p>
        </body>
        </html>
	`

	m.SetBody("text/html", htmlContent)

	// Melampirkan file
	for _, filePath := range filePaths {
		m.Attach(filePath)
	}

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("gagal mengirim email: %v", err)
	}
	return nil
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.UserRepo.FindByEmail(email)
}

func (s *UserService) GetUserByGoogleID(googleID string) (*models.User, error) {
	return s.UserRepo.FindByGoogleID(googleID)
}
