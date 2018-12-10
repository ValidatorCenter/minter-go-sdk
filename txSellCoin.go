package mintersdk

import (
	b64 "encoding/base64"
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Продажи монет
type TxSellCoinData struct {
	CoinToSell  string
	CoinToBuy   string
	ValueToSell float32
	// Other
	Payload string
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Продажи монет
func (c *SDK) TxSellCoin(t *TxSellCoinData) (string, error) {
	coinBuy := getStrCoin(t.CoinToBuy)
	coinSell := getStrCoin(t.CoinToSell)
	value := bip2pip_f64(float64(t.ValueToSell))
	coinGas := getStrCoin(t.GasCoin)
	valueGas := big.NewInt(t.GasPrice)

	privateKey, err := h2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	payComment := ""
	if t.Payload != "" {
		payComment = b64.StdEncoding.EncodeToString([]byte(t.Payload))
	}

	data := tr.SellCoinData{
		CoinToSell:  coinSell,
		ValueToSell: value,
		CoinToBuy:   coinBuy,
	}

	encodedData, err := serializeData(data)
	if err != nil {
		return "", err
	}

	_, nowNonce, err := c.Address(c.AccAddress)
	if err != nil {
		return "", err
	}

	tx := tr.Transaction{
		Nonce:         uint64(nowNonce + 1),
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeSellCoin,
		Data:          encodedData,
		Payload:       payComment,
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
