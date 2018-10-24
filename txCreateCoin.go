package mintersdk

import (
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Создания монеты
type TxCreateCoinData struct {
	Name                 string
	Symbol               string
	InitialAmount        int64
	InitialReserve       int64
	ConstantReserveRatio uint
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Создание монеты
func (c *SDK) TxCreateCoin(t *TxCreateCoinData) (string, error) {
	toCreate := GetStrCoin(t.Symbol)
	reserve := Bip2Pip_i64(t.InitialReserve)
	amount := Bip2Pip_i64(t.InitialAmount)
	coinGas := GetStrCoin(t.GasCoin)
	valueGas := big.NewInt(t.GasPrice)

	privateKey, err := H2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	data := tr.CreateCoinData{
		Name:                 t.Name,
		Symbol:               toCreate,
		InitialAmount:        amount,
		InitialReserve:       reserve,
		ConstantReserveRatio: t.ConstantReserveRatio,
	}

	encodedData, err := SerializeData(data)
	if err != nil {
		return "", err
	}

	tx := tr.Transaction{
		Nonce:         uint64(c.GetNonce(c.AccAddress) + 1),
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeCreateCoin,
		Data:          encodedData,
		SignatureType: tr.SigTypeSingle,
	}

	if err := tx.Sign(privateKey); err != nil {
		return "", err
	}

	resHash, err := c.SetTransaction(&tx)
	if err != nil {
		return "", err
	}
	return resHash, nil
}
