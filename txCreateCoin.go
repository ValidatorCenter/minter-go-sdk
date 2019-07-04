package mintersdk

import (
	"encoding/hex"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
	"github.com/MinterTeam/minter-go-node/core/types"
)

// Структура данных для Создания монеты
type TxCreateCoinData struct {
	Name                 string
	Symbol               string
	InitialAmount        int64
	InitialReserve       int64
	ConstantReserveRatio uint
	// Other
	Payload string
	// Gas
	GasCoin  string
	GasPrice int64
}

func (c *SDK) TxCreateCoinRLP(t *TxCreateCoinData) (string, error) {
	toCreate := getStrCoin(t.Symbol)
	reserve := bip2pip_i64(t.InitialReserve)
	amount := bip2pip_i64(t.InitialAmount)
	coinGas := getStrCoin(t.GasCoin)
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

	data := tr.CreateCoinData{
		Name:                 t.Name,
		Symbol:               toCreate,
		InitialAmount:        amount,
		InitialReserve:       reserve,
		ConstantReserveRatio: t.ConstantReserveRatio,
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
		Type:          tr.TypeCreateCoin,
		Data:          encodedData,
		Payload:       []byte(t.Payload),
		SignatureType: tr.SigTypeSingle,
	}

	if err := tx.Sign(privateKey); err != nil {
		return "", err
	}

	encodedTx, err := tx.Serialize()
	if err != nil {
		return "", err
	}

	strTxRPL := hex.EncodeToString(encodedTx)

	strRlpEnc := string(strTxRPL)

	return strRlpEnc, err
}

// Транзакция - Создание монеты
func (c *SDK) TxCreateCoin(t *TxCreateCoinData) (string, error) {
	strRlpEnc, err := c.TxCreateCoinRLP(t)
	if err != nil {
		return "", err
	}

	resHash, err := c.SetTransaction(strRlpEnc)
	if err != nil {
		return "", err
	}
	return resHash, nil
}
