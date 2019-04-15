package mintersdk

import (
	tr "github.com/MinterTeam/minter-go-node/core/transaction"
	"github.com/MinterTeam/minter-go-node/core/types"
)

// Структура данных для Декларирования мастерноды в кандидаты
type TxDeclareCandidacyData struct {
	PubKey     string // брать или с http://locallhost:3000 или с файла в conf/
	Commission uint
	Coin       string
	Stake      int64
	// Other
	Payload string
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Декларирования мастерноды в кандидаты
func (c *SDK) TxDeclareCandidacy(t *TxDeclareCandidacyData) (string, error) {
	myAddrss := getStrAddress(c.AccAddress)
	coin := getStrCoin(t.Coin)
	value := bip2pip_i64(t.Stake)
	pubkey := publicKey2Byte(t.PubKey)

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

	data := tr.DeclareCandidacyData{
		Address:    myAddrss,
		PubKey:     pubkey,
		Commission: t.Commission,
		Coin:       coin,
		Stake:      value,
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
		Type:          tr.TypeDeclareCandidacy,
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
