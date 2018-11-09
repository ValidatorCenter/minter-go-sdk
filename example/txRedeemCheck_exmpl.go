package main

import (
	"fmt"
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
	sdk "github.com/ValidatorCenter/minter-go-sdk"
)

func main() {
	sdk.SetAddressMn("https://minter-node-1.testnet.minter.network")

	AccAddress := "Mx..."
	AccPrivateKey := "..."
	passphrase := "password"

	coin := sdk.GetBaseCoin()
	privateKey, err := sdk.H2ECDSA(AccPrivateKey) // FIXME: rename H2ECDSA->h2ECDSA
	if err != nil {
		panic(err)
	}

	rawCheck, proof, err := sdk.CreateCheck(passphrase, 10, coin.String(), privateKey)
	if err != nil {
		panic(err)
	}

	data := tr.RedeemCheckData{
		RawCheck: rawCheck,
		Proof:    proof,
	}
	fmt.Println(data.String())

	encodedData, err := sdk.SerializeData(data)
	if err != nil {
		panic(err)
	}

	tx := tr.Transaction{
		Nonce:         uint64(sdk.GetNonce(AccAddress) + 1),
		GasPrice:      big.NewInt(1),
		GasCoin:       coin,
		Type:          tr.TypeRedeemCheck,
		Data:          encodedData,
		SignatureType: tr.SigTypeSingle,
	}

	if err := tx.Sign(privateKey); err != nil {
		panic(err)
	}

	resHash, err := sdk.SetTransaction(&tx)
	if err != nil {
		panic(err)
	}
	fmt.Println(resHash)
}
