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

	sndDt := m.TxSendCoinData{
		Coin:      "MNT",
		ToAddress: "Mxe64baa7d71c72e6914566b79ac361d139be22dc7", //Кому переводим
		Value:     10,
		GasCoin:   "MNT",
		GasPrice:  Gas,
	}

	resHash, err := sdk.TxSendCoin(&sndDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)

}
