package arweave

import (
	"arweave-datafeed/configs"
	"arweave-datafeed/utils"
	"arweave-datafeed/utils/log"
	"context"
	"math"
	"math/big"
)

// Transfer on arweave blockchain. Returns transaction hash, error
func Transfer(rates []byte, tag string, configuration *configs.ViperConfiguration) (string, error) {

	ar, err := NewTransactor(configuration.Get("nodeURL"))
	if err != nil {
		return "", err
	}

	arWallet := NewWallet()
	err = arWallet.LoadKeyFromFile(configuration.Get("walletFile"))
	if err != nil {
		return "", err
	}

	// display the balance of the wallet
	showBalance(arWallet, configuration)

	log.Printf("creating a transaction with a payload of %d bytes", len(rates))

	txBuilder, err := ar.CreateTransaction(context.Background(), tag, arWallet, "0", rates, "")
	if err != nil {
		return "", err
	}

	// sign the transaction
	txn, err := txBuilder.Sign(arWallet)
	if err != nil {
		return "", err
	}

	// send the transaction
	resp, err := ar.SendTransaction(context.Background(), txn)
	if err != nil {
		return "", err
	}

	log.Printf("arweave node responded %s", resp)

	return txn.Hash(), nil
}

// maybe this should be moved someplace else...
func showBalance(wallet *Wallet, configuration *configs.ViperConfiguration) {
	// display the balance of the wallet
	output, _, err := utils.GetRequest(configuration.Get("nodeURL") + "/wallet/" + wallet.Address() + "/balance")
	if err != nil {
		log.Error(err)
	}

	fBalance := new(big.Float)
	fBalance.SetString(string(output))
	arBalance := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(12)))

	log.Printf("Wallet Balance %s winston | %s AR", string(output), arBalance.String())
}
