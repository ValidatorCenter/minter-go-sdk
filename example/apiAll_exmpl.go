package main

import (
	"fmt"

	m "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	// Инициализация параметров
	sdk := m.SDK{
		MnAddress: "https://minter-node-1.testnet.minter.network:8841",
	}

	fmt.Println("##  1/11 ##")
	blnc, lastNmb, err := sdk.GetAddress("Mxdc7fcc63930bf81ebdce12b3bcef57b93e99a157")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Balance= %#v\nNonce=%d\n", blnc, lastNmb)

	fmt.Println("##  2/11 ##")
	blk, err := sdk.GetBlock(199)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetBlock= %#v\n", blk)

	fmt.Println("##  3/11 ##")
	cnd1, err := sdk.GetCandidate("Mp09f3548f7f4fc38ad2d0d8f805ec2cc1e35696012f95b8c6f2749e304a91efa2")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetCandidate= %#v\n", cnd1)

	fmt.Println("##  4/11 ##")
	cndAll, err := sdk.GetCandidates()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetCandidates= %#v\n", cndAll)

	fmt.Println("##  5/11 ##")
	coinInf, err := sdk.GetCoinInfo("VALIDATOR")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetCoinInfo= %#v\n", coinInf)

	fmt.Println("##  6/11 ##")
	stMn, err := sdk.GetStatus()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetStatus= %#v\n", stMn)

	fmt.Println("##  7/11 ##")
	trns, err := sdk.GetTransaction("Mtde93a53e774fe64f122274704896538fbf42c92ac8f8bd0dd307c8328600eff6")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetTransaction= %#v\n", trns)

	fmt.Println("## 8/11 ##")
	vldr, err := sdk.GetValidators()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetValidators= %#v\n", vldr)

	fmt.Println("## 9/11 ##")
	eCB, err := sdk.EstimateCoinBuy("MNT", "VALIDATOR", 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("EstimateCoinBuy= %#v\n", eCB)

	fmt.Println("## 10/11 ##")
	eCS, err := sdk.EstimateCoinSell("MNT", "VALIDATOR", 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("EstimateCoinSell= %#v\n", eCS)

	fmt.Println("## 11/11 ##")
	minGas, err := sdk.GetMinGas()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetMinGas= %#v\n", minGas)

	fmt.Println("##  Ok!  ##")
}
