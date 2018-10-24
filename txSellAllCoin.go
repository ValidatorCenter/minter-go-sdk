package mintersdk

import (
	"math/big"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Структура данных для Продажи всех монет
type TxSellAllCoinData struct {
	CoinToSell string
	CoinToBuy  string
	// Gas
	GasCoin  string
	GasPrice int64
}

// Транзакция - Продажи всех монет
func (c *SDK) TxSellAllCoin(t *TxSellAllCoinData) (string, error) {
	coinBuy := GetStrCoin(t.CoinToBuy)
	coinSell := GetStrCoin(t.CoinToSell)

	coinGas := GetStrCoin(t.GasCoin)
	valueGas := big.NewInt(t.GasPrice)

	privateKey, err := H2ECDSA(c.AccPrivateKey)
	if err != nil {
		return "", err
	}

	data := tr.SellAllCoinData{
		CoinToSell: coinSell,
		CoinToBuy:  coinBuy,
	}

	encodedData, err := SerializeData(data)
	if err != nil {
		return "", err
	}

	tx := tr.Transaction{
		Nonce:         uint64(c.GetNonce(c.AccAddress) + 1),
		GasPrice:      valueGas,
		GasCoin:       coinGas,
		Type:          tr.TypeSellAllCoin,
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
