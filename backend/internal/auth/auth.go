// Package auth handles authentication using Google OAuth
package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "randomString"
	MaxAge = 60 * 60 * 24 * 30 // 30 days
	isProd = false             // Controls cookie security: false=dev (HTTP allowed), true=prod (HTTPS only)
)

// NewAuth initializes the authentication system
func NewAuth() {
	// Load environment variables from .env file
	fmt.Println("started loading")
	err := godotenv.Load()
	// if .env file is not found, log an error
	if err != nil {
		fmt.Println("loading failed")
		log.Fatal("Error loading .env file")
	}

	// Get environment variables from .env file
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	fmt.Println("enviroment variables loaded successfully")

	// Create a new cookie store
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	fmt.Println("cookie store created successfully")

	gothic.Store = store

	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, "http://localhost:3000/auth/google/callback"),
	)
	fmt.Println("goth providers created successfully")
}
