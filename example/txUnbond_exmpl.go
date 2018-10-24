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

	unbDt := m.TxUnbondData{
		PubKey:   "Mp5c87d35a7adb055f54140ba03c0eed418ddc7c52ff7a63fc37a0e85611388610",
		Coin:     "MNT",
		Value:    10,
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxUnbond(&unbDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
