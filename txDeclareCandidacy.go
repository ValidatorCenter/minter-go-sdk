package mintersdk

import (
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Декларирования мастерноды в кандидаты
type TxDeclareCandidacyData struct {
	PubKey     string // брать или с http://locallhost:3000 или с файла в conf/
	Commission uint
	Coin       string
	Stake      int64
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
	valueGas := big.NewInt(t.GasPrice)
	privateKey, err := h2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
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

	tx := tr.Transaction{
		Nonce:         uint64(c.GetNonce(c.AccAddress) + 1),
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeDeclareCandidacy,
		Data:          encodedData,
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
