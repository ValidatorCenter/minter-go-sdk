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
	Payload   string
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

	nowNonce, err := c.GetNonce(c.AccAddress)
	if err != nil {
		return "", err
	}

	/*
	   Nonce - int, used for prevent transaction reply.
	   Gas Price - big int, used for managing transaction fees.
	   Gas Coin - 10 bytes, symbol of a coin to pay fee
	   Type - type of transaction (see below).
	   Data - data of transaction (depends on transaction type).
	   Payload (arbitrary bytes) - arbitrary user-defined bytes.
	   Service Data - reserved field.
	   Signature Type - single or multisig transaction.
	   Signature Data - digital signature of transaction.
	*/
	// TODO: b64 "encoding/base64" Payload шифрование сообщения

	tx := tr.Transaction{
		Nonce:         uint64(nowNonce + 1),
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeSend,
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
