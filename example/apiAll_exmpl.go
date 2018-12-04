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

	fmt.Println("##  1/12 ##")
	lastNmb, err := sdk.GetNonce("Mxdc7fcc63930bf81ebdce12b3bcef57b93e99a157")
	if err != nil {
		panic(err)
	}
	fmt.Println("GetNonce=", lastNmb)

	fmt.Println("##  2/12 ##")
	blnc, err := sdk.GetBalance("Mxdc7fcc63930bf81ebdce12b3bcef57b93e99a157")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetBalance= %#v\n", blnc)

	fmt.Println("##  3/12 ##")
	vlm, err := sdk.GetBaseCoinVolume(1)
	if err != nil {
		panic(err)
	}
	fmt.Println("GetBaseCoinVolume=", vlm)

	fmt.Println("##  4/12 ##")
	blk, err := sdk.GetBlock(199)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetBlock= %#v\n", blk)

	fmt.Println("##  5/12 ##")
	cnd1, err := sdk.GetCandidate("Mp5c87d35a7adb055f54140ba03c0eed418ddc7c52ff7a63fc37a0e85611388610")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetCandidate= %#v\n", cnd1)

	fmt.Println("##  6/12 ##")
	cndAll, err := sdk.GetCandidates()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetCandidates= %#v\n", cndAll)

	fmt.Println("##  7/12 ##")
	coinInf, err := sdk.GetCoinInfo("ABCDEF23")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetCoinInfo= %#v\n", coinInf)

	fmt.Println("##  8/12 ##")
	stMn, err := sdk.GetStatus()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetStatus= %#v\n", stMn)

	fmt.Println("##  9/12 ##")
	trns, err := sdk.GetTransaction("Mt8e091163d410fb34c621ed3f30b38192b36de836")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetTransaction= %#v\n", trns)

	fmt.Println("## 10/12 ##")
	vldr, err := sdk.GetValidators()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetValidators= %#v\n", vldr)

	fmt.Println("## 11/12 ##")
	eCB, err := sdk.EstimateCoinBuy("MNT", "VALIDATOR", 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("EstimateCoinBuy= %#v\n", eCB)

	fmt.Println("## 12/12 ##")
	eCS, err := sdk.EstimateCoinSell("MNT", "VALIDATOR", 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("EstimateCoinSell= %#v\n", eCS)

	fmt.Println("##  Ok!  ##")
}
