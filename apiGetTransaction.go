package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

//curl -s 'localhost:8841/api/transaction/{hash}'
type node_transaction struct {
	Code   int
	Result TransResponse
}

type TransResponse struct {
	Hash     string            `json:"hash"`
	RawTx    string            `json:"raw_tx"`
	Height   int64             `json:"height"`
	Index    int64             `json:"index"`
	From     string            `json:"from"`
	Nonce    uint64            `json:"nonce"`
	GasPrice *big.Int          `json:"gas_price"`
	GasCoin  string            `json:"gas_coin"`
	TxResult ResponseDeliverTx `json:"tx_result"`
	Type     byte              `json:"type"`
	Data     TransData         `json:"data"`
	Payload  []byte            `json:"payload"`

	Tags TagKeyValue2 `json:"tags"`
}

type ResponseDeliverTx struct {
	GasWanted int64        `json:"gas_wanted"`
	GasUsed   int64        `json:"gas_used"`
	Tags      TagKeyValue2 `json:"tags" bson:"tags"`
}

type TagKeyValue2 struct {
	TxCoinToBuy  string `json:"tx.coin_to_buy"`
	TxCoinToSell string `json:"tx.coin_to_sell"`
	TxFrom       string `json:"tx.from"`
	TxReturn     string `json:"tx.return"`
	TxSellAmount string `json:"tx.sell_amount"`
	//tx.type	"\u0002"
}

type TransData struct {
	//*** type1 - TYPE_SEND
	To string `json:"to"`
	//Coin   string `json:"coin"`
	//Value  string `json:"value"`
	//*** type2 - TYPE_SELL_COIN
	ValueToSell string `json:"value_to_sell"`
	//CoinToSell string `json:"coin_to_sell"`
	//CoinToBuy  string `json:"coin_to_buy"`
	//*** type3 - TYPE_SELL_ALL_COIN
	//CoinToSell string `json:"coin_to_sell"`
	//CoinToBuy  string `json:"coin_to_buy"`
	//*** type4 - TYPE_BUY_COIN
	CoinToBuy  string `json:"coin_to_buy"`
	ValueToBuy string `json:"value_to_buy"`
	CoinToSell string `json:"coin_to_sell"`
	//*** type5 - TYPE_CREATE_COIN
	Name                 string `json:"name"`                   // название монеты
	CoinSymbol           string `json:"coin_symbol"`            // символ монеты
	InitialAmount        string `json:"initial_amount"`         //  Amount of coins to issue. Issued coins will be available to sender account.
	InitialReserve       string `json:"initial_reserve"`        // Initial reserve in base coin.
	ConstantReserveRatio int    `json:"constant_reserve_ratio"` // uint, should be from 10 to 100 (в %).
	//*** type6 - TYPE_DECLARE_CANDIDACY
	Address    string `json:"address"`
	Commission int    `json:"commission"`
	//Stake	string `json:"stake"`
	//PubKey  string `json:"pub_key"`
	//Coin    string `json:"coin"`
	//*** type7 - TYPE_DELEGATE
	Stake string `json:"stake"`
	//PubKey  string `json:"pub_key"`
	//Coin    string `json:"coin"`
	//*** type8 - TYPE_UNBOUND
	PubKey string `json:"pub_key"`
	Coin   string `json:"coin"`
	Value  string `json:"value"`
	//*** type9 - TYPE_REDEEM_CHECK
	//*** type10 - TYPE_SET_CANDIDATE_ONLINE
	//*** type11 - TYPE_SET_CANDIDATE_OFFLINE
}

// получаем содержимое транзакции по её хэшу
func (c *SDK) GetTransaction(hash string) TransResponse {
	url := fmt.Sprintf("%s/api/transaction/%s", c.MnAddress, hash)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data node_transaction
	json.Unmarshal(body, &data)
	return data.Result
}
