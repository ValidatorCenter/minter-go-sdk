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
	lastNmb := sdk.GetNonce("Mxdc7fcc63930bf81ebdce12b3bcef57b93e99a157")
	fmt.Println("GetNonce=", lastNmb)

	fmt.Println("##  2/12 ##")
	blnc := sdk.GetBalance("Mxdc7fcc63930bf81ebdce12b3bcef57b93e99a157")
	fmt.Printf("GetBalance= %#v\n", blnc)

	fmt.Println("##  3/12 ##")
	vlm := sdk.GetBaseCoinVolume(1)
	fmt.Println("GetBaseCoinVolume=", vlm)

	fmt.Println("##  4/12 ##")
	blk := sdk.GetBlock(199)
	fmt.Printf("GetBlock= %#v\n", blk)

	fmt.Println("##  5/12 ##")
	cnd1 := sdk.GetCandidate("Mp5c87d35a7adb055f54140ba03c0eed418ddc7c52ff7a63fc37a0e85611388610")
	fmt.Printf("GetCandidate= %#v\n", cnd1)

	fmt.Println("##  6/12 ##")
	cndAll := sdk.GetCandidates()
	fmt.Printf("GetCandidates= %#v\n", cndAll)

	fmt.Println("##  7/12 ##")
	coinInf := sdk.GetCoinInfo("ABCDEF23")
	fmt.Printf("GetCoinInfo= %#v\n", coinInf)

	fmt.Println("##  8/12 ##")
	stMn := sdk.GetStatus()
	fmt.Printf("GetStatus= %#v\n", stMn)

	fmt.Println("##  9/12 ##")
	trns := sdk.GetTransaction("Mt8e091163d410fb34c621ed3f30b38192b36de836")
	fmt.Printf("GetTransaction= %#v\n", trns)

	fmt.Println("## 10/12 ##")
	vldr := sdk.GetValidators()
	fmt.Printf("GetValidators= %#v\n", vldr)

	fmt.Println("## 11/12 ##")
	eCB := sdk.EstimateCoinBuy("MNT", "VALIDATOR", 1)
	fmt.Printf("EstimateCoinBuy= %#v\n", eCB)

	fmt.Println("## 12/12 ##")
	eCS := sdk.EstimateCoinSell("MNT", "VALIDATOR", 1)
	fmt.Printf("EstimateCoinSell= %#v\n", eCS)

	fmt.Println("##  Ok!  ##")
}
