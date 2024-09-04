// Di file internal/app/routes/routes.go, kita akan mendefinisikan rute untuk register dan login.

package routes

import (
	"BACKEND-GOLANG-MVVM/internal/app/controllers"

	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userController *controllers.UserController,
	notificationController *controllers.NotificationController,
	authController *controllers.AuthController,
	baseController *controllers.BaseController) *gin.Engine {
	router := gin.Default()

	//Apply CORS middleware
	// router.Use(middleware.CORS())
	appEnv := os.Getenv("APP_ENV")

	// Tentukan path static berdasarkan environment
	var staticPath string
	if appEnv == "docker" {
		staticPath = "/app/static"
	} else {
		staticPath = "./static"
	}

	router.Static("/static", "./static")

	// Serve the index.html file
	router.StaticFile("/", staticPath+"/index.html")
	router.StaticFile("/register", staticPath+"/register.html")
	router.StaticFile("/notification", staticPath+"/notification.html")
	router.StaticFile("/formula", staticPath+"/formula.html")
	router.StaticFile("/dashboard", staticPath+"/dashboard.html")
	router.StaticFile("/firebase-messaging-sw.js", staticPath+"/firebase-messaging-sw.js")

	api := router.Group("/api")
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)
		api.DELETE("/logout/:userID", userController.Logout)
		api.POST("/notify", notificationController.Notify)

		// Rute untuk Google OAuth Login & Register
		api.GET("/google-login", authController.GoogleLogin)
		api.GET("/google-register", authController.GoogleRegister)

		// Rute untuk Google OAuth Register
		api.GET("/callback", authController.GoogleCallback) // Satu rute callback

		api.POST("/calculate-expression", baseController.CalculateExpression) // Mengganti dari /calculate-salary
		api.GET("/functions", baseController.GetAvailableFunctions)           // Tetap sama

		api.POST("/send-to-queue", userController.SendToQueue)
		api.GET("/last-message", userController.GetLastMessage)
	}

	// Socket.IO routes
	// router.GET("/socket.io/*any", gin.WrapH(services.Server))
	// router.POST("/socket.io/*any", gin.WrapH(services.Server))

	return router
}
