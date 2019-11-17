package main

import (
	"arweave-datafeed/utils/log"
	"fmt"
	"github.com/robfig/cron/v3"
)

// startSync uploads exchange rates to Arweave
func startSync() {

	log.Println("~~~ SYNC STARTED ~~~")

	c := cron.New()
	_, err := c.AddFunc("* * * * *", uploadRatesToArweave)
	if err != nil {
		log.Error(err)
	}
	c.Start()

	// forever
	//for {
	//

	//
	//	lastError = ""
	//	fmt.Println("----------------- ZzZzZzZz -----------------")
	//
	//}

}

func uploadRatesToArweave() {

	rates, err := getExchangeRates()
	if err != nil {
		lastError = err.Error()
		log.Error(err)
	}

	fmt.Printf("rates %s", string(rates))
	//timeTag := time.Now().UTC().Format("2006-01-02T15")
}
