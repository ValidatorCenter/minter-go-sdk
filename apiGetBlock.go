package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Содержимое блока
type node_block struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  BlockResponse
	Error   ErrorStruct
}

type ErrorStruct struct {
	Code    int
	Message string
	Data    string
}

type BlockResponse struct {
	Hash          string                    `json:"hash" bson:"hash" gorm:"hash"`
	HeightTx      string                    `json:"height" bson:"-" gorm:"-"`
	Height        int                       `json:"height_i32" bson:"height_i32" gorm:"height_i32"`
	Time          time.Time                 `json:"time" bson:"time" gorm:"time"`
	NumTxsTx      string                    `json:"num_txs" bson:"-" gorm:"-"`
	NumTxs        int                       `json:"num_txs_i32" bson:"num_txs_i32" gorm:"num_txs_i32"`
	TotalTxsTx    string                    `json:"total_txs" bson:"-" gorm:"-"`
	TotalTxs      int                       `json:"total_txs_i32" bson:"total_txs_i32" gorm:"total_txs_i32"`
	Transactions  []TransResponse           `json:"transactions" bson:"transactions" gorm:"transactions"`
	BlockRewardTx string                    `json:"block_reward" bson:"-" gorm:"-"`
	BlockReward   float32                   `json:"block_reward_f32" bson:"block_reward_f32" gorm:"block_reward_f32"`
	SizeTx        string                    `json:"size" bson:"-" gorm:"-"`
	Size          int                       `json:"size_i32" bson:"size_i32" gorm:"size_i32"`
	Validators    []BlockValidatorsResponse `json:"validators" bson:"validators" gorm:"validators"`
	Proposer      string                    `json:"proposer" bson:"proposer" gorm:"proposer"` // PubKey пропозер блока
}

type BlockValidatorsResponse struct {
	PubKey string `json:"pubkey" bson:"pubkey" gorm:"pubkey"`
	Signed bool   `json:"signed" bson:"signed" gorm:"signed"` // подписал-true, или пропустил false
}

// type TransResponse struct --- в apiGetTransaction.go
// type TransData struct --- в apiGetTransaction.go

// получаем содержимое блока по его ID
func (c *SDK) GetBlock(id int) (BlockResponse, error) {
	url := fmt.Sprintf("%s/block?height=%d", c.MnAddress, id)
	res, err := http.Get(url)
	if err != nil {
		return BlockResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return BlockResponse{}, err
	}

	var data node_block
	json.Unmarshal(body, &data)

	data.Result.BlockReward = pipStr2bip_f32(data.Result.BlockRewardTx) // вознаграждение за блок

	data.Result.Height, err = strconv.Atoi(data.Result.HeightTx)
	if err != nil {
		return BlockResponse{}, err
	}

	data.Result.NumTxs, err = strconv.Atoi(data.Result.NumTxsTx)
	if err != nil {
		return BlockResponse{}, err
	}

	data.Result.Size, err = strconv.Atoi(data.Result.SizeTx)
	if err != nil {
		return BlockResponse{}, err
	}

	for iStep, _ := range data.Result.Transactions {
		data.Result.Transactions[iStep].HeightTx = data.Result.HeightTx

		err = manipulationTransaction(c, &data.Result.Transactions[iStep])
		if err != nil {
			return BlockResponse{}, err
		}
	}

	return data.Result, nil
}
