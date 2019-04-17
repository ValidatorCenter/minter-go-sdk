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

	declDt := m.TxDeclareCandidacyData{
		PubKey:     "Mp09f3548f7f4fc38ad2d0d8f805ec2cc1e35696012f95b8c6f2749e304a91efa2", // "Mp....",
		Commission: 10,                                                                   // 10%
		Coin:       "MNT",
		Stake:      1000,
		// Gas
		GasCoin:  "MNT",
		GasPrice: Gas,
	}

	resHash, err := sdk.TxDeclareCandidacy(&declDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)
}
