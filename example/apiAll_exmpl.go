package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	// Инициализация параметров
	sdk := m.SDK{
		MnAddress: "https://minter-node-1.testnet.minter.network",
	}

	fmt.Println("##  1/10 ##")
	blnc, lastNmb, err := sdk.GetAddress("Mxdc7fcc63930bf81ebdce12b3bcef57b93e99a157")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Balance= %#v\nNonce=%d\n", blnc, lastNmb)

	fmt.Println("##  2/10 ##")
	blk, err := sdk.GetBlock(199)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetBlock= %#v\n", blk)

	fmt.Println("##  3/10 ##")
	cnd1, err := sdk.GetCandidate("Mp7555e8a7b2fea6e0c45dd92075338076dc330a6b3e09130720d9946a4421a97a")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetCandidate= %#v\n", cnd1)

	fmt.Println("##  4/10 ##")
	cndAll, err := sdk.GetCandidates()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetCandidates= %#v\n", cndAll)

	fmt.Println("##  5/10 ##")
	coinInf, err := sdk.GetCoinInfo("VALIDATOR")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetCoinInfo= %#v\n", coinInf)

	fmt.Println("##  6/10 ##")
	stMn, err := sdk.GetStatus()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetStatus= %#v\n", stMn)

	fmt.Println("##  7/10 ##")
	trns, err := sdk.GetTransaction("Mt2220a952f18d3573d1d7c2c59f9969365a88da6a5d29ea960bcb2a17c97c30cc")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetTransaction= %#v\n", trns)

	fmt.Println("## 8/10 ##")
	vldr, err := sdk.GetValidators()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetValidators= %#v\n", vldr)

	fmt.Println("## 9/10 ##")
	eCB, err := sdk.EstimateCoinBuy("MNT", "VALIDATOR", 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("EstimateCoinBuy= %#v\n", eCB)

	fmt.Println("## 10/10 ##")
	eCS, err := sdk.EstimateCoinSell("MNT", "VALIDATOR", 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("EstimateCoinSell= %#v\n", eCS)

	fmt.Println("##  Ok!  ##")
}
