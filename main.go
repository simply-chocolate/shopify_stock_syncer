package main

import (
	"fmt"
	"log"
	"time"

	"pcn_stock_syncer/utils"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Printf("%v: Started the Script \n", time.Now().Format("2006-01-02 15:04:05"))

	err = utils.HandleSyncStock()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v: Success \n", time.Now().Format("2006-01-02 15:04:05"))

	fmt.Printf("%v: Started the Cron Scheduler", time.Now().UTC().Format("2006-01-02 15:04:05"))

	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Cron("0 8,10,12,14,16 * * *").SingletonMode().Do(func() {
		fmt.Printf("%v: Started the Script \n", time.Now().Format("2006-01-02 15:04:05"))

		err = utils.HandleSyncStock()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v: Success \n", time.Now().Format("2006-01-02 15:04:05"))

	})

	s.StartBlocking()

}
