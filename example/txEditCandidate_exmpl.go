package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccPrivateKey: "...",
		ChainMainnet:  false,
	}

	Gas, _ := sdk.GetMinGas()

	edtCDt := m.TxEditCandidateData{
		PubKey:        "Mp09f3548f7f4fc38ad2d0d8f805ec2cc1e35696012f95b8c6f2749e304a91efa2",
		RewardAddress: "Mx7a86fb0d770062decdca3dc5fed15800d5a65000",
		OwnerAddress:  "Mx58a1441883708813ba546345a0ed0ce765f1dad1",
		// Gas
		GasCoin:  "MNT",
		GasPrice: Gas,
	}

	resHash, err := sdk.TxEditCandidate(&edtCDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)
}
