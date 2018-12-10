package mintersdk

import (
	b64 "encoding/base64"
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Передачи монет
type TxSendCoinData struct {
	Coin      string
	ToAddress string
	Value     float32
	// Other
	Payload string
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Передача монет
func (c *SDK) TxSendCoin(t *TxSendCoinData) (string, error) {

	to := getStrAddress(t.ToAddress)
	coin := getStrCoin(t.Coin)
	value := bip2pip_f64(float64(t.Value))

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

	data := tr.SendData{
		Coin:  coin,
		To:    to,
		Value: value,
	}

	encodedData, err := serializeData(data)
	if err != nil {
		return "", err
	}

	_, nowNonce, err := c.GetAddress(c.AccAddress)
	if err != nil {
		return "", err
	}

	tx := tr.Transaction{
		Nonce:         uint64(nowNonce + 1),
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeSend,
		Data:          encodedData,
		Payload:       []byte(payComment),
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
