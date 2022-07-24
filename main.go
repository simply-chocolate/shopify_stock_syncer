package main

import (
	"fmt"
	"log"

	"pcn_stock_syncer/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = utils.HandleSyncStock()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("I'm gonna sync some stock one day.")
}
