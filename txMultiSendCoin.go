package mintersdk

import (
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
		List: listAddrs, // max=100 транзакций за 1 раз
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
		Type:          tr.TypeMultisend,
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
