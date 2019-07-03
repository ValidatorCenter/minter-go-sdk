package mintersdk

import (
	tr "github.com/MinterTeam/minter-go-node/core/transaction"
	"github.com/MinterTeam/minter-go-node/core/types"
)

// Структура данных для Продажи всех монет
type TxSellAllCoinData struct {
	CoinToSell string
	CoinToBuy  string
	// Other
	Payload string
	// Gas
	GasCoin  string
	GasPrice int64
}

func (c *SDK) TxSellAllCoinRLP(t *TxSellAllCoinData) (string, error) {
	coinBuy := getStrCoin(t.CoinToBuy)
	coinSell := getStrCoin(t.CoinToSell)

	coinGas := getStrCoin(t.GasCoin)
	valueGas := uint32(t.GasPrice)

	privateKey, err := h2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	data := tr.SellAllCoinData{
		CoinToSell: coinSell,
		CoinToBuy:  coinBuy,
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
		Type:          tr.TypeSellAllCoin,
		Data:          encodedData,
		Payload:       []byte(t.Payload),
		SignatureType: tr.SigTypeSingle,
	}

	if err := tx.Sign(privateKey); err != nil {
		return "", err
	}

	encodedTx, err := tx.Serialize()
	if err != nil {
		fmt.Println("ERROR: TxSellAllCoin::tx.Serialize")
		return "", err
	}

	strTxRPL := hex.EncodeToString(encodedTx)

	strRlpEnc := string(strTxRPL)

	return strRlpEnc, err

}

// Транзакция - Продажи всех монет
func (c *SDK) TxSellAllCoin(t *TxSellAllCoinData) (string, error) {
	strRlpEnc, err := TxSellAllCoinRLP(t)
	if err != nil {
		return "", err
	}

	resHash, err := c.SetTransaction(strRlpEnc)
	if err != nil {
		return "", err
	}
	return resHash, nil
}
