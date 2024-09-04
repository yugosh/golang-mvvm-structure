package services

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

var Server *socketio.Server
var ReceivedMessages []string

func InitSocketIO() {
	Server = socketio.NewServer(nil)

	Server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("Connected:", s.ID())
		return nil
	})

	Server.OnEvent("/", "request_last_message", func(s socketio.Conn, msg string) {
		if len(ReceivedMessages) > 0 {
			s.Emit("last_message", ReceivedMessages[len(ReceivedMessages)-1])
		}
	})

	Server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("Disconnected:", s.ID(), "Reason:", reason)
	})

	go Server.Serve()
}

func BroadcastMessage(message string) {
	log.Printf("Broadcasting message: %s", message)
	// ReceivedMessages = append(ReceivedMessages, message)
	ReceivedMessages = append(ReceivedMessages, message)
	if Server != nil {
		log.Println("Socket.IO Server is initialized")
		Server.BroadcastToRoom("/", "", "new_message", message)
	} else {
		log.Println("Socket.IO Server is not initialized")
	}
}
