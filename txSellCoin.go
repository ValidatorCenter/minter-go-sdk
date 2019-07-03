package mintersdk

import (
	tr "github.com/MinterTeam/minter-go-node/core/transaction"
	"github.com/MinterTeam/minter-go-node/core/types"
)

// Структура данных для Продажи монет
type TxSellCoinData struct {
	CoinToSell  string
	CoinToBuy   string
	ValueToSell float32
	// Other
	Payload string
	// Gas
	GasCoin  string
	GasPrice int64
}

func (c *SDK) TxSellCoinRLP(t *TxSellCoinData) (string, error) {
	coinBuy := getStrCoin(t.CoinToBuy)
	coinSell := getStrCoin(t.CoinToSell)
	value := bip2pip_f64(float64(t.ValueToSell))
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

	data := tr.SellCoinData{
		CoinToSell:  coinSell,
		ValueToSell: value,
		CoinToBuy:   coinBuy,
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
		Type:          tr.TypeSellCoin,
		Data:          encodedData,
		Payload:       []byte(t.Payload),
		SignatureType: tr.SigTypeSingle,
	}

	if err := tx.Sign(privateKey); err != nil {
		return "", err
	}

	encodedTx, err := tx.Serialize()
	if err != nil {
		fmt.Println("ERROR: TxSellCoin::tx.Serialize")
		return "", err
	}

	strTxRPL := hex.EncodeToString(encodedTx)

	strRlpEnc := string(strTxRPL)

	return strRlpEnc, err

}

// Транзакция - Продажи монет
func (c *SDK) TxSellCoin(t *TxSellCoinData) (string, error) {
	strRlpEnc, err := TxSellCoinRLP(t)
	if err != nil {
		return "", err
	}

	resHash, err := c.SetTransaction(strRlpEnc)
	if err != nil {
		return "", err
	}
	return resHash, nil
}
