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

	chDt := m.TxCreateCkeckData{
		Coin:     "MNT",
		Stake:    10,
		Password: "pswrd123",
		Nonce:    102,
	}

	resCheck, err := sdk.TxCreateCheck(&chDt)
	if err != nil {
		panic(err)
	}
	fmt.Println(resCheck)

	// Redeem
	rchDt := m.TxRedeemCheckData{
		Check:    resCheck,
		Password: "pswrd123",
		GasCoin:  "MNT",
		GasPrice: Gas,
	}

	resHash, err := sdk.TxRedeemCheck(&rchDt)
	if err != nil {
		panic(err)
	}
	fmt.Println("HashTx:", resHash)

}
