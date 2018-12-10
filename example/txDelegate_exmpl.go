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
		PubKey:   "Mp7555e8a7b2fea6e0c45dd92075338076dc330a6b3e09130720d9946a4421a97a",
		Stake:    10,
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxDelegate(&delegDt)
	if err != nil {
		panic(err)
	}
	fmt.Println("HashTx:", resHash)

}
