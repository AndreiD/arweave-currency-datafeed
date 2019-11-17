package main

import (
	"arweave-datafeed/arweave"
	"arweave-datafeed/configs"
	"arweave-datafeed/utils"
	"arweave-datafeed/utils/log"
	"fmt"
	"github.com/robfig/cron/v3"
	"math"
	"math/big"
	"time"
)

// startCronService starts the cron service every hour
func startCronService() {

	c := cron.New()
	// replace first with * from every minute, and with 0 for every hour
	_, err := c.AddFunc("0 * * * *", uploadRatesToArweave)
	if err != nil {
		log.Error(err)
	}
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

	showBalance(configuration)

	log.Printf("transfer to Arweave finished in %s successfully. Tx ID %s", fmt.Sprintf("%s", time.Since(start)), txID)
}

// maybe this should be moved someplace else...
func showBalance(configuration *configs.ViperConfiguration) {

	arWallet := arweave.NewWallet()
	err := arWallet.LoadKeyFromFile(configuration.Get("walletFile"))
	if err != nil {
		log.Error(err)
	}

	output, _, err := utils.GetRequest(configuration.Get("nodeURL") + "/wallet/" + arWallet.Address() + "/balance")
	if err != nil {
		log.Error(err)
	}

	fBalance := new(big.Float)
	fBalance.SetString(string(output))
	arBalance := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(12)))

	log.Printf("Wallet Balance %s winston | %s AR", string(output), arBalance.String())
}
