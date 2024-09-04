package controllers

import (
	"BACKEND-GOLANG-MVVM/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	// Jika ada dependency yang perlu disuntikkan, bisa ditambahkan di sini
}

func NewNotificationController() *NotificationController {
	return &NotificationController{}
}

func (nc *NotificationController) Notify(ctx *gin.Context) {
	token := ctx.Query("token") // Mengambil parameter dari query string
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	title := "Notification Title"
	body := "This is the body of the notification"

	err := services.SendNotification(token, title, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Notification sent successfully"})
}
