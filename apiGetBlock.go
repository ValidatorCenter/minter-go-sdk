package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
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
	Precommits   []BlockPrecommitResponse   `json:"precommits"`
	BlockReward  string                     `json:"block_reward"`
}

type BlockPrecommitResponse struct {
	ValidatorAddress string `json:"validator_address"`
	ValidatorIndex   string `json:"validator_index"`
	Height           string `json:"height"`
	Round            string `json:"round"`
	Timestamp        string `json:"timestamp"`
	Type             int    `json:"type"`
	Signature        string `json:"signature"`
	/*block_id {
	string hash	"E3822B41EFE536BA5FA6FCC832B7D21D6B1638B7"
	parts {
		string total	"1"
		string hash	"9103B3C507E59A9C34EA470E057F8ED96766CF10"
		}
	}*/

}

type BlockTransactionResponse struct {
	Hash        string            `json:"hash"`
	RawTx       string            `json:"raw_tx"`
	From        string            `json:"from"`
	Nonce       uint64            `json:"nonce"`
	GasPrice    *big.Int          `json:"gas_price"`
	Type        byte              `json:"type"`
	Data        TransactionData   `json:"data"`
	Payload     []byte            `json:"payload"`
	ServiceData []byte            `json:"service_data"`
	Gas         int64             `json:"gas"`
	GasCoin     CoinSymbol        `json:"gas_coin"`
	TxResult    ResponseDeliverTx `json:"tx_result"`
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
