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

	slAllDt := m.TxSellAllCoinData{
		CoinToSell: "MNT",
		CoinToBuy:  "ABCDEF24",
		GasCoin:    "MNT",
		GasPrice:   Gas,
	}

	resHash, err := sdk.TxSellAllCoin(&slAllDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
