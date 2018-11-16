package mintersdk

import (
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Статус кандидата вкл/выкл
type TxSetCandidateData struct {
	PubKey   string
	Activate bool // Вкл./выкл мастерноду
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Вкл./выкл мастерноду в валидаторы
func (c *SDK) TxSetCandidate(t *TxSetCandidateData) (string, error) {
	pubkey := publicKey2Byte(t.PubKey)

	coinGas := getStrCoin(t.GasCoin)
	valueGas := big.NewInt(t.GasPrice)

	privateKey, err := h2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	var typeTx byte
	var data interface{}

	if t.Activate == true {
		data = tr.SetCandidateOnData{
			PubKey: pubkey,
		}
		typeTx = tr.TypeSetCandidateOnline
	} else {
		data = tr.SetCandidateOffData{
			PubKey: pubkey,
		}
		typeTx = tr.TypeSetCandidateOffline
	}

	encodedData, err := serializeData(data)
	if err != nil {
		return "", err
	}

	tx := tr.Transaction{
		Nonce:         uint64(c.GetNonce(c.AccAddress) + 1),
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          typeTx,
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
