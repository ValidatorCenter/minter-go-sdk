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

	slAllDt := m.TxSellAllCoinData{
		CoinToSell: "MNT",
		CoinToBuy:  "ABCDEF24",
		GasCoin:    "MNT",
		GasPrice:   1,
	}

	resHash, err := sdk.TxSellAllCoin(&slAllDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
