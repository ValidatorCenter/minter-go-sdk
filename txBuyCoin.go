package mintersdk

import (
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Покупки монет
type TxBuyCoinData struct {
	CoinToSell string
	CoinToBuy  string
	ValueToBuy int64
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Купить монету
func (c *SDK) TxBuyCoin(t *TxBuyCoinData) (string, error) {
	coinBuy := GetStrCoin(t.CoinToBuy)
	coinSell := GetStrCoin(t.CoinToSell)
	coinGas := GetStrCoin(t.GasCoin)
	value := Bip2Pip_i64(t.ValueToBuy)
	valueGas := big.NewInt(t.GasPrice)

	privateKey, err := H2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	data := tr.BuyCoinData{
		CoinToSell: coinSell,
		ValueToBuy: value,
		CoinToBuy:  coinBuy,
	}

	encodedData, err := SerializeData(data)
	if err != nil {
		return "", err
	}

	tx := tr.Transaction{
		Nonce:         uint64(c.GetNonce(c.AccAddress) + 1),
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeBuyCoin,
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
