package mintersdk

import (
	b64 "encoding/base64"
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Передачи монет нескольком адресатам
type TxMultiSendCoinData struct {
	List []TxOneSendCoinData
	// Other
	Payload string
	// Gas
	GasCoin  string
	GasPrice int64
}

type TxOneSendCoinData struct {
	Coin      string
	ToAddress string
	Value     float32
}

// Транзакция - Передача монет нескольким адресатам
func (c *SDK) TxMultiSendCoin(t *TxMultiSendCoinData) (string, error) {
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

	listAddrs := []tr.MultisendDataItem{}

	for _, dtSend := range t.List {
		to := getStrAddress(dtSend.ToAddress)
		coin := getStrCoin(dtSend.Coin)
		value := bip2pip_f64(float64(dtSend.Value))
		listAddrs = append(listAddrs, tr.MultisendDataItem{
			Coin:  coin,
			To:    to,
			Value: value,
		})
	}

	data := tr.MultisendData{
		List: listAddrs,
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
		Type:          tr.TypeMultisend,
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
