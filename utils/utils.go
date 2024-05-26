package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func DevLog(log string) {
	if os.Getenv("DEVMODE") == "true" {
		fmt.Printf("%v: %v\n", time.Now().Format("2006-01-02 15:04:05"), log)
	}
}

func LoadEnv() {
	godotenv.Load()

	if os.Getenv("SHOPIFY_ADDRESS") == "" {
		panic("Error loading environment variable SHOPIFY_ADDRESS")
	}
	if os.Getenv("SHOPIFY_API_VERSION") == "" {
		panic("Error loading environment variable SHOPIFY_API_VERSION")
	}
	if os.Getenv("SHOPIFY_KEY") == "" {
		panic("Error loading environment variable SHOPIFY_KEY")
	}
	if os.Getenv("SHOPIFY_PASS") == "" {
		panic("Error loading environment variable SHOPIFY_PASS")
	}
	if os.Getenv("PCN_ADRESS") == "" {
		panic("Error loading environment variable PCN_ADRESS")
	}
	if os.Getenv("PCN_OLSUSER") == "" {
		panic("Error loading environment variable PCN_OLSUSER")
	}
	if os.Getenv("PCN_CID") == "" {
		panic("Error loading environment variable PCN_CID")
	}
	if os.Getenv("PCN_OLSPASS") == "" {
		panic("Error loading environment variable PCN_OLSPASS")
	}
	if os.Getenv("PCN_AUTH_UN") == "" {
		panic("Error loading environment variable PCN_AUTH_UN")
	}
	if os.Getenv("PCN_AUTH_PW") == "" {
		panic("Error loading environment variable PCN_AUTH_PW")
	}
	if os.Getenv("TEAMS_WEBHOOK_URL") == "" {
		panic("Error loading environment variable TEAMS_WEBHOOK_URL")
	}
	if os.Getenv("DEVMODE") == "" {
		panic("Error loading environment variable DEVMODE")
	}
}
