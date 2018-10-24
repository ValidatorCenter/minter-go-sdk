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

	creatDt := m.TxCreateCoinData{
		Name:                 "Test coin 24",
		Symbol:               "ABCDEF24",
		InitialAmount:        100,
		InitialReserve:       100,
		ConstantReserveRatio: 50,
		// Gas
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxCreateCoin(&creatDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
