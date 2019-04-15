package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccPrivateKey: "...",
	}

	Gas, _ := sdk.GetMinGas()

	unbDt := m.TxUnbondData{
		PubKey:   "Mp504815c4a47418aa37b17248e359cb5a5272bd8f416eb9d1d3b8ba95b394296f",
		Coin:     "MNT",
		Value:    10,
		GasCoin:  "MNT",
		GasPrice: Gas,
	}

	resHash, err := sdk.TxUnbond(&unbDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
