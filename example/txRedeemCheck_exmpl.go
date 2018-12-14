package main

import (
	"fmt"
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
	m "github.com/ValidatorCenter/minter-go-sdk"

	"crypto/ecdsa"

	"github.com/MinterTeam/minter-go-node/core/types"
	"github.com/MinterTeam/minter-go-node/crypto"
	"github.com/MinterTeam/minter-go-node/rlp"
)

func _h2ECDSA(AccPrivateKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(AccPrivateKey)
}

// Преобразует строку в монету
func _getStrCoin(coin string) types.CoinSymbol {
	var mntV types.CoinSymbol
	copy(mntV[:], []byte(coin))
	return mntV
}

func _serializeData(data interface{}) ([]byte, error) {
	return rlp.EncodeToBytes(data)
}

func main() {
	var err error
	sdk := m.SDK{
		MnAddress:     "https://minter-node-1.testnet.minter.network",
		AccPrivateKey: "...",
	}
	passphrase := "password"

	sdk.AccAddress, err = m.GetAddressPrivateKey(sdk.AccPrivateKey)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(sdk.AccAddress)
	}

	coin := m.GetBaseCoin()
	privateKey, err := _h2ECDSA(sdk.AccPrivateKey)
	if err != nil {
		panic(err)
	}

	//FIXME: Проблема в toolsFuncCheck.go
	rawCheck, proof, err := m.CreateCheck(passphrase, 10, coin, privateKey)
	if err != nil {
		panic(err)
	}

	data := tr.RedeemCheckData{
		RawCheck: rawCheck,
		Proof:    proof,
	}
	fmt.Println(data.String())

	encodedData, err := _serializeData(data)
	if err != nil {
		panic(err)
	}

	_, nowNonce, err := sdk.GetAddress(sdk.AccAddress)
	if err != nil {
		panic(err)
	}

	tx := tr.Transaction{
		Nonce:         uint64(nowNonce + 1),
		GasPrice:      big.NewInt(1),
		GasCoin:       _getStrCoin(coin),
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
