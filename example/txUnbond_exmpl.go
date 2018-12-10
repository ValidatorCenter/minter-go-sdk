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
		PubKey:   "Mp7555e8a7b2fea6e0c45dd92075338076dc330a6b3e09130720d9946a4421a97a",
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
