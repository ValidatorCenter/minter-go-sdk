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
	Hash         string                     `json:"hash"`
	Height       int64                      `json:"height"`
	Time         time.Time                  `json:"time"`
	NumTxs       int64                      `json:"num_txs"`
	TotalTxs     int64                      `json:"total_txs"`
	Transactions []BlockTransactionResponse `json:"transactions"`
	Events       []BlockEventsResponse      `json:"events"`
	Precommits   []BlockPrecommitResponse   `json:"precommits"`
	BlockReward  string                     `json:"block_reward"`
	Size         int                        `json:"size"`
}

type BlockPrecommitResponse struct {
	ValidatorAddress string      `json:"validator_address"`
	ValidatorIndex   string      `json:"validator_index"`
	Height           string      `json:"height"`
	Round            string      `json:"round"`
	Timestamp        string      `json:"timestamp"`
	Type             int         `json:"type"`
	Signature        string      `json:"signature"`
	BlockID          BlockIDData `json:"block_id"`
}

type BlockIDData struct {
	Hash  string    `json:"hash"`
	Parts PartsData `json:"parts"`
}

type PartsData struct {
	Total string `json:"total"`
	Hash  string `json:"hash"`
}

type BlockEventsResponse struct {
	Type  string         `json:"type"`
	Value EventValueData `json:"value"`
}

type EventValueData struct {
	Role            string `json:"role"`
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ValidatorPubKey string `json:"validator_pub_key"`
}

type BlockTransactionResponse struct {
	Hash        string          `json:"hash"`
	RawTx       string          `json:"raw_tx"`
	From        string          `json:"from"`
	Nonce       int             `json:"nonce"`
	GasPrice    int             `json:"gas_price"`
	Type        int             `json:"type"`
	Data        TransactionData `json:"data"`
	Payload     []byte          `json:"payload"`
	ServiceData []byte          `json:"service_data"`
	Gas         int             `json:"gas"`
	GasCoin     string          `json:"gas_coin"`
	GasUsed     int             `json:"gas_used"`
	//TxResult    ResponseDeliverTx `json:"tx_result"` // TODO: del
	//Tags    TagKeyValue2 `json:"tags"` // TODO: нет необходимости в нём
}

type TransactionData struct {
	Coin  string `json:"coin"`
	To    string `json:"to"`
	Value string `json:"value"`
}

// type ResponseDeliverTx struct --- в apiGetTransaction.go

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
