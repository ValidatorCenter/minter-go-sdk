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

	sndDt := m.TxSetCandidateData{
		PubKey:   "Mp2891198c692c351bc55ac60a03c82649fa920f7ad20bd290a0c4e774e916e9de",
		Activate: true, //true-"on", false-"off"
		GasCoin:  "MNT",
		GasPrice: Gas,
	}

	resHash, err := sdk.TxSetCandidate(&sndDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
