package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// TODO: add analytics
// TODO: use typing feature - cafe havana is for.. foodies. people who add to
// 			 the conversation. folks who have made history. and you!

func main() {
	// Setup logger.
	logger := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	// Load .env file.
	if os.Getenv("ENVIRONMENT") != "production" {
		if err := godotenv.Load(); err != nil {
			panic("could not load env!")
		}
	}

	// Minify assets.
	if os.Getenv("ENVIRONMENT") != "production" {
		if success := minifyAssets(); !success {
			panic("could not minify assets!")
		}
	}

	// Build handler.
	h := handler{Logger: logger}

	// Setup routes.
	router := buildRoutes(h)

	// Run server.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		panic("could not start server!")
	}
}
