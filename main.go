package main

import (
	"fmt"
	"log"
	"pcn_stock_syncer/pcn_api_wrapper"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	result, err := pcn_api_wrapper.PcnApiGetStockData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("I'm gonna sync some stock one day.")
}
