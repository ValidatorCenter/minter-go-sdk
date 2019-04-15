package mintersdk

import (
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
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

// Транзакция - Продажи всех монет
func (c *SDK) TxSellAllCoin(t *TxSellAllCoinData) (string, error) {
	coinBuy := getStrCoin(t.CoinToBuy)
	coinSell := getStrCoin(t.CoinToSell)

	coinGas := getStrCoin(t.GasCoin)
	valueGas := big.NewInt(t.GasPrice)

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
		Type:          tr.TypeSellAllCoin,
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
