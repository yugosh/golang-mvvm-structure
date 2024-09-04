package main

import (
	"BACKEND-GOLANG-MVVM/internal/app/controllers"
	"BACKEND-GOLANG-MVVM/internal/app/middleware"
	"BACKEND-GOLANG-MVVM/internal/app/repositories"
	"BACKEND-GOLANG-MVVM/internal/app/routes"
	"BACKEND-GOLANG-MVVM/internal/app/services"
	"BACKEND-GOLANG-MVVM/internal/config"
	"BACKEND-GOLANG-MVVM/internal/db"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// services.InitSocketIO()

	// Initialize the rabbitmq
	config.InitRabbitMQ()
	defer config.CloseRabbitMQ()

	go services.StartConsumer("test_queue")

	// Initialize the redis
	config.InitRedis()

	// Initialize the database
	database, err := db.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected successfully:", database)

	// Initialize repositories
	userRepo := repositories.NewUserRepository(database)

	// Initialize services
	userService := services.NewUserService(userRepo)
	formulaService := services.NewFormulaService()
	expressionService := services.NewExpressionService(formulaService)

	// Initialize controllers
	userController := controllers.NewUserController(userService)
	notificationController := controllers.NewNotificationController()
	baseController := controllers.NewBaseController(expressionService, formulaService)

	authController := &controllers.AuthController{
		UserService: *userService,
	}

	// Allowed origins
	allowedOrigins := []string{
		"http://localhost:8080",
		"https://2cc1-116-90-182-149.ngrok-free.app/",
		"google.com",
		// Tambahkan origins lain jika diperlukan, misalnya:
		// "http://localhost:3000",
	}

	// Setup Gin router
	router := gin.Default()

	// Apply CORS middleware dengan allowed origins
	router.Use(middleware.CORS(allowedOrigins))

	// Setup routes
	router = routes.SetupRouter(userController, notificationController, authController, baseController)

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
