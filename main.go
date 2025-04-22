package main

import (
	"fmt"
	"github.com/joho/godotenv"
	//"github.com/slack-go/slack"
	"log"
	"os"
)

func main() {
	// Try to get the Slack token from environment variables
	slackToken := os.Getenv("slackToken")

	// If not found, attempt to load from .env file
	if slackToken == "" {
		log.Println("No Slack token found in system environments")

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
			return
		}
		slackToken = os.Getenv("slackToken")
		if slackToken == "" {
			log.Fatal("Slack token taken from .env file")
		} else {
			fmt.Println("No Slack token found in system environments nor .env file")
		}

	}
	// Token successfully found
	fmt.Println("Slack token loaded successfully:", slackToken)

}
