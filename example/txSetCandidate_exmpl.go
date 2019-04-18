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
		PubKey:   "Mp09f3548f7f4fc38ad2d0d8f805ec2cc1e35696012f95b8c6f2749e304a91efa2",
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
