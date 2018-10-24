package mintersdk

import (
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Отзыва монет
type TxUnbondData struct {
	PubKey string
	Coin   string
	Value  int64
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Отозвать монеты из делегированных в валидатора
func (c *SDK) TxUnbond(t *TxUnbondData) (string, error) {

	pubkey := PublicKey2Byte(t.PubKey)
	coin := GetStrCoin(t.Coin)
	value := Bip2Pip_i64(t.Value)

	coinGas := GetStrCoin(t.GasCoin)
	valueGas := big.NewInt(t.GasPrice)

	privateKey, err := H2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	data := tr.UnbondData{
		PubKey: pubkey,
		Coin:   coin,
		Value:  value,
	}

	encodedData, err := SerializeData(data)
	if err != nil {
		return "", err
	}

	tx := tr.Transaction{
		Nonce:         uint64(c.GetNonce(c.AccAddress) + 1),
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeUnbond,
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
