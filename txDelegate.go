package mintersdk

import (
	b64 "encoding/base64"
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Делегирования
type TxDelegateData struct {
	PubKey string
	Coin   string
	Stake  float32
	// Other
	Payload string
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Делегирование
func (c *SDK) TxDelegate(t *TxDelegateData) (string, error) {
	coin := getStrCoin(t.Coin)
	coinGas := getStrCoin(t.GasCoin)
	value := bip2pip_f64(float64(t.Stake))
	valueGas := big.NewInt(t.GasPrice)
	pubkey := publicKey2Byte(t.PubKey)
	privateKey, err := h2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	payComment := ""
	if t.Payload != "" {
		payComment = b64.StdEncoding.EncodeToString([]byte(t.Payload))
	}

	data := tr.DelegateData{
		PubKey: pubkey,
		Coin:   coin,
		Stake:  value,
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
		Type:          tr.TypeDelegate,
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
