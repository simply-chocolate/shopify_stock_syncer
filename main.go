package main

import (
	"fmt"
	"log"
	"time"

	"pcn_stock_syncer/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Printf("%v Started the Script \n", time.Now().Format("2006-01-02 15:04:05"))

	err = utils.HandleSyncStock()
	if err != nil {
		log.Fatal(err)
		fmt.Printf("%v", err)
	}

	fmt.Printf("%v Success \n", time.Now().Format("2006-01-02 15:04:05"))

}
