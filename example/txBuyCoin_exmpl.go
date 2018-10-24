package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccAddress:    "Mx...",
		AccPrivateKey: "...",
	}

	buyDt := m.TxBuyCoinData{
		CoinToSell: "ABCDEF23",
		CoinToBuy:  "MNT",
		ValueToBuy: 1,
		// Gas
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxBuyCoin(&buyDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
