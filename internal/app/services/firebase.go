package services

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitFirebase() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("config/serviceAccountKey.json") // ganti dengan path ke service account key kamu
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	FirebaseApp = app
}

func SendNotification(token string, title string, body string) error {
	ctx := context.Background()
	client, err := FirebaseApp.Messaging(ctx)
	if err != nil {
		return err
	}

	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	_, err = client.Send(ctx, message)
	return err
}
