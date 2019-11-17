package main

import (
	"arweave-datafeed/arweave"
	"arweave-datafeed/configs"
	"arweave-datafeed/utils"
	"github.com/gin-gonic/gin"
	"math"
	"math/big"
	"net/http"
	"runtime"
)

// shows the last error...
var lastError string

// shows if the bot is running
func healthHandler(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	c.JSON(http.StatusOK, gin.H{"status": "alive", "last_error": lastError, "alloc": bToMb(m.Alloc), "total_alloc": bToMb(m.TotalAlloc),
		"sys": bToMb(m.Sys), "num_gc": m.NumGC})
}

// converts bytes to Mb
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// balance of tokens of the wallet in use
func balanceHandler(c *gin.Context) {

	configuration, ok := c.MustGet("configuration").(*configs.ViperConfiguration)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get the configuration"})
		return
	}

	// get my address
	arWallet := arweave.NewWallet()
	err := arWallet.LoadKeyFromFile(configuration.Get("walletFile"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output, _, err := utils.GetRequest(configuration.Get("nodeURL") + "/wallet/" + arWallet.Address() + "/balance")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fBalance := new(big.Float)
	fBalance.SetString(string(output))
	arBalance := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(12)))

	c.JSON(http.StatusOK, gin.H{"winston": string(output), "ar": arBalance.String()})
}
