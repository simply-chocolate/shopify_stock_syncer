package main

import (
	"fmt"
	"log"
	"pcn_stock_syncer/shopify_api_wrapper"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	/*
		result, err := pcn_api_wrapper.PcnApiGetStockData()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)
	*/
	result1, err := shopify_api_wrapper.ShopifyApiGetInventoryId(shopify_api_wrapper.ShopifyApiQueryParams{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", result1.Body.Products[5])

	fmt.Println("I'm gonna sync some stock one day.")
}
