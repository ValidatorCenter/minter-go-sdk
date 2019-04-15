package mintersdk

import (
	"encoding/hex"
	"fmt"
	"strings"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
	"github.com/MinterTeam/minter-go-node/core/types"
)

// Структура данных для Создания монеты
type TxCreateCkeckData struct {
	Coin     string
	Stake    float32
	Password string
	Nonce    uint64
}

// НЕ!Транзакция - Создание чека
func (c *SDK) TxCreateCheck(t *TxCreateCkeckData) (string, error) {
	//coin := getStrCoin(t.Coin)
	privateKey, err := h2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	// СОЗДАНИЕ ЧЕКА
	rawCheck, err := createCheck(t.Password, t.Stake, t.Coin, privateKey, t.Nonce)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Mc%s", string(hex.EncodeToString(rawCheck))), nil
}

type TxRedeemCheckData struct {
	Check    string
	Password string
	// Other
	Payload string
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Погашение чека (обналичивание)
func (c *SDK) TxRedeemCheck(t *TxRedeemCheckData) (string, error) {
	coinGas := getStrCoin(t.GasCoin)
	valueGas := uint32(t.GasPrice)

	rawCheck := ""
	// убираем Mc
	if t.Check[0:2] == "Mc" {
		rawCheck = strings.TrimLeft(t.Check, "Mc")
	} else {
		rawCheck = t.Check
	}
	rawCheckHex, err := hex.DecodeString(rawCheck)
	if err != nil {
		return "", err
	}

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

	// ОБНАЛИЧИВАЕМ СЕБЕ
	proof, err := checkCashingProof(t.Password, privateKey)
	if err != nil {
		return "", err
	}

	data := tr.RedeemCheckData{
		RawCheck: rawCheckHex,
		Proof:    proof,
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
		Type:          tr.TypeRedeemCheck,
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
