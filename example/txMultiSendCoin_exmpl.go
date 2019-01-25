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

	cntList := []m.TxOneSendCoinData{}

	// Первый адресат
	cntList = append(cntList, m.TxOneSendCoinData{
		Coin:      "MNT",
		ToAddress: "Mxe64baa7d71c72e6914566b79ac361d139be22dc7", //Кому переводим
		Value:     10,
	})

	// Второй адресат
	cntList = append(cntList, m.TxOneSendCoinData{
		Coin:      "MNT",
		ToAddress: "Mxe64baa7d71c72e6914566b79ac361d139be22dc7", //Кому переводим
		Value:     10,
	})

	mSndDt := m.TxMultiSendCoinData{
		List:     cntList,
		GasCoin:  "MNT",
		GasPrice: 1,
	}

	resHash, err := sdk.TxMultiSendCoin(&mSndDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
