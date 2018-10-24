package mintersdk

import (
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Делегирования
type TxDelegateData struct {
	PubKey string
	Coin   string
	Stake  int64
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Делегирование
func (c *SDK) TxDelegate(t *TxDelegateData) (string, error) {
	coin := GetStrCoin(t.Coin)
	coinGas := GetStrCoin(t.GasCoin)
	value := Bip2Pip_i64(t.Stake)
	valueGas := big.NewInt(t.GasPrice)
	pubkey := PublicKey2Byte(t.PubKey)
	privateKey, err := H2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	data := tr.DelegateData{
		PubKey: pubkey,
		Coin:   coin,
		Stake:  value,
	}

	encodedData, err := SerializeData(data)
	if err != nil {
		return "", err
	}

	tx := tr.Transaction{
		Nonce:         uint64(c.GetNonce(c.AccAddress) + 1),
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeDelegate,
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
