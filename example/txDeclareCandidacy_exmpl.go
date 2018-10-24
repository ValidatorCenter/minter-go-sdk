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

	declDt := m.TxDeclareCandidacyData{
		PubKey:     "Mp2891198c692c351bc55ac60a03c82649fa920f7ad20bd290a0c4e774e916e9de", // "Mp....",
		Commission: 10,                                                                   // 10%
		Coin:       "MNT",
		Stake:      100,
		// Gas
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxDeclareCandidacy(&declDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)
}
