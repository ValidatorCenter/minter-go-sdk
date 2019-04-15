package mintersdk

import (
	tr "github.com/MinterTeam/minter-go-node/core/transaction"
	"github.com/MinterTeam/minter-go-node/core/types"
)

// Структура данных для Покупки монет
type TxBuyCoinData struct {
	CoinToSell string
	CoinToBuy  string
	ValueToBuy float32
	// Other
	Payload string
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Купить монету
func (c *SDK) TxBuyCoin(t *TxBuyCoinData) (string, error) {
	coinBuy := getStrCoin(t.CoinToBuy)
	coinSell := getStrCoin(t.CoinToSell)
	coinGas := getStrCoin(t.GasCoin)
	value := bip2pip_f64(float64(t.ValueToBuy))
	valueGas := uint32(t.GasPrice)

	privateKey, err := h2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	if c.AccAddress == "" {
		c.AccAddress, err = GetAddressPrivateKey(c.AccPrivateKey)
		if err != nil {
			return "", err
		}
	}

	data := tr.BuyCoinData{
		CoinToSell: coinSell,
		ValueToBuy: value,
		CoinToBuy:  coinBuy,
	}

	encodedData, err := serializeData(data)
	if err != nil {
		return "", err
	}

	_, nowNonce, err := c.GetAddress(c.AccAddress)
	if err != nil {
		return "", err
	}

	var _ChainID types.ChainID
	if c.ChainMainnet {
		_ChainID = types.ChainMainnet
	} else {
		_ChainID = types.ChainTestnet
	}

	tx := tr.Transaction{
		Nonce:         uint64(nowNonce + 1),
		ChainID:       _ChainID,
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeBuyCoin,
		Data:          encodedData,
		Payload:       []byte(t.Payload),
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
