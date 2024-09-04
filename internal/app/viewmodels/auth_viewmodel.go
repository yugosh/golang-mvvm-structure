package viewmodels

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig *oauth2.Config
var OauthStateString string

func init() {
	// Memuat variabel lingkungan dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/callback", // Sesuaikan dengan URL callback Anda
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),        // Client ID dari Google Cloud Console
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),    // Client Secret dari Google Cloud Console
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	OauthStateString = "dj8MNF!9d.zSje2" // Gunakan string random untuk state

	// Debugging: memastikan bahwa variabel lingkungan sudah di-load
	log.Println("GOOGLE_CLIENT_ID:", os.Getenv("GOOGLE_CLIENT_ID"))
	log.Println("GOOGLE_CLIENT_SECRET:", os.Getenv("GOOGLE_CLIENT_SECRET"))
}

// https://accounts.google.com/o/oauth2/auth?client_id=163130246348-710dp7v8dufhogsjc344d1gvigvq0hu3.apps.googleusercontent.com&redirect_uri=http://localhost:8080/api/callback&response_type=code&scope=https://www.googleapis.com/auth/userinfo.email&state=dj8MNF!9d.zSje2
// https://accounts.google.com/o/oauth2/auth?client_id=163130246348-710dp7v8dufhogsjc344d1gvigvq0hu3.apps.googleusercontent.com&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fapi%2Fcallback&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email&state=dj8MNF%219d.zSje2
