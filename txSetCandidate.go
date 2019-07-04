package mintersdk

import (
	"encoding/hex"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
	"github.com/MinterTeam/minter-go-node/core/types"
)

// Структура данных для Статус кандидата вкл/выкл
type TxSetCandidateData struct {
	PubKey   string
	Activate bool // Вкл./выкл мастерноду
	// Other
	Payload string
	// Gas
	GasCoin  string
	GasPrice int64
}

func (c *SDK) TxSetCandidateRLP(t *TxSetCandidateData) (string, error) {
	pubkey := publicKey2Byte(t.PubKey)

	coinGas := getStrCoin(t.GasCoin)
	valueGas := uint32(t.GasPrice)

	privateKey, err := h2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	var typeTx tr.TxType
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
		Type:          typeTx,
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

// Транзакция - Вкл./выкл мастерноду в валидаторы
func (c *SDK) TxSetCandidate(t *TxSetCandidateData) (string, error) {
	strRlpEnc, err := c.TxSetCandidateRLP(t)
	if err != nil {
		return "", err
	}

	resHash, err := c.SetTransaction(strRlpEnc)
	if err != nil {
		return "", err
	}
	return resHash, nil
}
