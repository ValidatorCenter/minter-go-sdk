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

	delegDt := m.TxDelegateData{
		Coin:     "MNT",
		PubKey:   "Mp5c87d35a7adb055f54140ba03c0eed418ddc7c52ff7a63fc37a0e85611388610",
		Stake:    100,
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxDelegate(&delegDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
