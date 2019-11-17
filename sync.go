package main

import (
	"arweave-datafeed/arweave"
	"arweave-datafeed/utils/log"
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

// startCronService starts the cron service every hour
func startCronService() {
	uploadRatesToArweave()
	c := cron.New()
	//_, err := c.AddFunc("* * * * *", uploadRatesToArweave) // replace first with * from every minute, and with 0 for every hour
	//if err != nil {
	//	log.Error(err)
	//}
	c.Start()
}

// uploadRatesToArweave does just that
func uploadRatesToArweave() {

	start := time.Now()

	rates, err := getExchangeRates()
	if err != nil {
		lastError = err.Error()
		log.Error(err)
	}

	log.Printf("rates %s", string(rates))
	timeTag := time.Now().UTC().Format("2006-01-02T15")
	log.Printf("time tag %s", timeTag)

	txID, err := arweave.Transfer(rates, timeTag, configuration)
	if err != nil {
		lastError = err.Error()
		log.Error(err)
	}

	log.Printf("transfer to Arweave finished in %s successfully. Tx ID %s", fmt.Sprintf("%s", time.Since(start)), txID)
}
