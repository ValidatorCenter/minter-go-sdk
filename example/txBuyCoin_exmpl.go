package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccPrivateKey: "...",
		ChainMainnet:  false,
	}

	Gas, _ := sdk.GetMinGas()

	buyDt := m.TxBuyCoinData{
		CoinToSell: "ABCDEF23",
		CoinToBuy:  "MNT",
		ValueToBuy: 1,
		// Gas
		GasCoin:  "MNT",
		GasPrice: Gas,
	}

	resHash, err := sdk.TxBuyCoin(&buyDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
