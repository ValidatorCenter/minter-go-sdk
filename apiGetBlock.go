package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"math/big"
	"net/http"
	"time"
)

// Содержимое блока
type node_block struct {
	Code   int
	Result BlockResponse
}

type BlockResponse struct {
	Hash         string                     `json:"hash" bson:"hash" gorm:"hash"`
	Height       int64                      `json:"height" bson:"height" gorm:"height"`
	Time         time.Time                  `json:"time" bson:"time" gorm:"time"`
	NumTxs       int64                      `json:"num_txs" bson:"num_txs" gorm:"num_txs"`
	TotalTxs     int64                      `json:"total_txs" bson:"total_txs" gorm:"total_txs"`
	Transactions []BlockTransactionResponse `json:"transactions" bson:"transactions" gorm:"transactions"`
	Events       []BlockEventsResponse      `json:"events" bson:"events" gorm:"events"`
	Precommits   []BlockPrecommitResponse   `json:"precommits" bson:"precommits" gorm:"precommits"`
	BlockReward  string                     `json:"block_reward" bson:"block_reward" gorm:"block_reward"`
	Size         int                        `json:"size" bson:"size" gorm:"size"`
}

type BlockPrecommitResponse struct {
	ValidatorAddress string      `json:"validator_address" bson:"validator_address" gorm:"validator_address"`
	ValidatorIndex   string      `json:"validator_index" bson:"validator_index" gorm:"validator_index"`
	Height           string      `json:"height" bson:"height" gorm:"height"`
	Round            string      `json:"round" bson:"round" gorm:"round"`
	Timestamp        string      `json:"timestamp" bson:"timestamp" gorm:"timestamp"`
	Type             int         `json:"type" bson:"type" gorm:"type"`
	Signature        string      `json:"signature" bson:"signature" gorm:"signature"`
	BlockID          BlockIDData `json:"block_id" bson:"block_id" gorm:"block_id"`
}

type BlockIDData struct {
	Hash  string    `json:"hash" bson:"hash" gorm:"hash"`
	Parts PartsData `json:"parts" bson:"parts" gorm:"parts"`
}

type PartsData struct {
	Total string `json:"total" bson:"total" gorm:"total"`
	Hash  string `json:"hash" bson:"hash" gorm:"hash"`
}

type BlockEventsResponse struct {
	Type  string         `json:"type" bson:"type" gorm:"type"`
	Value EventValueData `json:"value" bson:"value" gorm:"value"`
}

type EventValueData struct {
	Role            string `json:"role" bson:"role" gorm:"role"`
	Address         string `json:"address" bson:"address" gorm:"address"`
	Amount          string `json:"amount" bson:"amount" gorm:"amount"`
	ValidatorPubKey string `json:"validator_pub_key" bson:"validator_pub_key" gorm:"validator_pub_key"`
}

type BlockTransactionResponse struct {
	Hash        string          `json:"hash" bson:"hash" gorm:"hash"`
	RawTx       string          `json:"raw_tx" bson:"raw_tx" gorm:"raw_tx"`
	From        string          `json:"from" bson:"from" gorm:"from"`
	Nonce       int             `json:"nonce" bson:"nonce" gorm:"nonce"`
	GasPrice    int             `json:"gas_price" bson:"gas_price" gorm:"gas_price"`
	Type        int             `json:"type" bson:"type" gorm:"type"`
	Data        TransactionData `json:"data" bson:"data" gorm:"data"`
	Payload     []byte          `json:"payload" bson:"payload" gorm:"payload"`
	ServiceData []byte          `json:"service_data" bson:"service_data" gorm:"service_data"`
	Gas         int             `json:"gas" bson:"gas" gorm:"gas"`
	GasCoin     string          `json:"gas_coin" bson:"gas_coin" gorm:"gas_coin"`
	GasUsed     int             `json:"gas_used" bson:"gas_used" gorm:"gas_used"`
	//TxResult    ResponseDeliverTx `json:"tx_result"` // TODO: del
	//Tags    TagKeyValue2 `json:"tags" bson:"tags" gorm:"tags"` // TODO: нет необходимости в нём
}

type TransactionData struct {
	Coin  string `json:"coin" bson:"coin" gorm:"coin"`
	To    string `json:"to" bson:"to" gorm:"to"`
	Value string `json:"value" bson:"value" gorm:"value"`
}

// type ResponseDeliverTx struct --- в apiGetTransaction.go // TODO: del

// получаем содержимое блока по его ID
func (c *SDK) GetBlock(id int) BlockResponse {
	url := fmt.Sprintf("%s/api/block/%d", c.MnAddress, id)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data node_block
	json.Unmarshal(body, &data)
	//fmt.Printf("%d\n", data.Result.Height)
	return data.Result
}
