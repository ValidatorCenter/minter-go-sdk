package mintersdk

import (
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

	if c.ChainMainnet {
		ChainID = ChainMainnet
	} else {
		ChainID = ChainTestnet
	}

	tx := tr.Transaction{
		Nonce:         uint64(nowNonce + 1),
		ChainID:       ChainID,
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeSend,
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
