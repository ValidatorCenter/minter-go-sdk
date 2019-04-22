package mintersdk

import (
	tr "github.com/MinterTeam/minter-go-node/core/transaction"
	"github.com/MinterTeam/minter-go-node/core/types"
)

// Структура данных для Статус кандидата вкл/выкл
type TxEditCandidateData struct {
	PubKey        string
	RewardAddress string
	OwnerAddress  string
	// Other
	Payload string
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Декларирования мастерноды в кандидаты
func (c *SDK) TxEditCandidate(t *TxEditCandidateData) (string, error) {
	oAddrss := getStrAddress(t.OwnerAddress)
	rAddrss := getStrAddress(t.RewardAddress)
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

	data := tr.EditCandidateData{
		PubKey:        pubkey,
		RewardAddress: rAddrss,
		OwnerAddress:  oAddrss,
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
		Type:          tr.TypeEditCandidate,
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
