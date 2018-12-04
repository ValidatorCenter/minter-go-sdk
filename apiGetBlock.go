package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Содержимое блока
type node_block struct {
	Code   int
	Result BlockResponse
}

type BlockResponse struct {
	Hash          string                     `json:"hash" bson:"hash" gorm:"hash"`
	Height        int64                      `json:"height" bson:"height" gorm:"height"`
	Time          time.Time                  `json:"time" bson:"time" gorm:"time"`
	NumTxs        int64                      `json:"num_txs" bson:"num_txs" gorm:"num_txs"`
	TotalTxs      int64                      `json:"total_txs" bson:"total_txs" gorm:"total_txs"`
	Transactions  []BlockTransactionResponse `json:"transactions" bson:"transactions" gorm:"transactions"`
	Events        []BlockEventsResponse      `json:"events" bson:"events" gorm:"events"`
	Precommits    []BlockPrecommitResponse   `json:"precommits" bson:"precommits" gorm:"precommits"`
	BlockRewardTx string                     `json:"block_reward" bson:"-" gorm:"-"`
	BlockReward   float32                    `json:"block_reward_f32" bson:"block_reward_f32" gorm:"block_reward_f32"`
	Size          int                        `json:"size" bson:"size" gorm:"size"`
}

type BlockPrecommitResponse struct {
	ValidatorAddress string      `json:"validator_address" bson:"validator_address" gorm:"validator_address"`
	ValidatorIndexTx string      `json:"validator_index" bson:"-" gorm:"-"`
	ValidatorIndex   int         `json:"validator_index_i32" bson:"validator_index_i32" gorm:"validator_index_i32"`
	HeightTx         string      `json:"height" bson:"-" gorm:"-"`
	Height           int         `json:"height_i32" bson:"height_i32" gorm:"height_i32"`
	RoundTx          string      `json:"round" bson:"-" gorm:"-"`
	Round            int         `json:"round_i32" bson:"round_i32" gorm:"round_i32"`
	Timestamp        time.Time   `json:"timestamp" bson:"timestamp" gorm:"timestamp"`
	Type             int         `json:"type" bson:"type" gorm:"type"`
	Signature        string      `json:"signature" bson:"signature" gorm:"signature"`
	BlockID          BlockIDData `json:"block_id" bson:"block_id" gorm:"block_id"`
}

type BlockIDData struct {
	Hash  string    `json:"hash" bson:"hash" gorm:"hash"`
	Parts PartsData `json:"parts" bson:"parts" gorm:"parts"`
}

type PartsData struct {
	TotalTx string `json:"total" bson:"" gorm:""`
	Total   int    `json:"total_i32" bson:"total_i32" gorm:"total_i32"`
	Hash    string `json:"hash" bson:"hash" gorm:"hash"`
}

type BlockEventsResponse struct {
	Type  string         `json:"type" bson:"type" gorm:"type"`
	Value EventValueData `json:"value" bson:"value" gorm:"value"`
}

type EventValueData struct {
	Role            string  `json:"role" bson:"role" gorm:"role"` //DAO,Developers,Validator,Delegator
	Address         string  `json:"address" bson:"address" gorm:"address"`
	AmountTx        string  `json:"amount" bson:"-" gorm:"-"`
	Amount          float32 `json:"amount_f32" bson:"amount_f32" gorm:"amount_f32"`
	ValidatorPubKey string  `json:"validator_pub_key" bson:"validator_pub_key" gorm:"validator_pub_key"`
}

type BlockTransactionResponse struct {
	Hash        string      `json:"hash" bson:"hash" gorm:"hash"`
	RawTx       string      `json:"raw_tx" bson:"raw_tx" gorm:"raw_tx"`
	From        string      `json:"from" bson:"from" gorm:"from"`
	Nonce       int         `json:"nonce" bson:"nonce" gorm:"nonce"`
	GasPrice    int         `json:"gas_price" bson:"gas_price" gorm:"gas_price"`
	Type        int         `json:"type" bson:"type" gorm:"type"`
	DataTx      TransData   `json:"data" bson:"-" gorm:"-"`
	Data        interface{} `json:"-" bson:"data" gorm:"data"`
	Payload     string      `json:"payload" bson:"payload" gorm:"payload"`
	ServiceData []byte      `json:"service_data" bson:"service_data" gorm:"service_data"`
	Gas         int         `json:"gas" bson:"gas" gorm:"gas"`
	GasCoin     string      `json:"gas_coin" bson:"gas_coin" gorm:"gas_coin"`
	GasUsed     int         `json:"gas_used" bson:"gas_used" gorm:"gas_used"`
	//TxResult    ResponseDeliverTx `json:"tx_result"` // TODO: del
	//Tags    TagKeyValue2 `json:"tags" bson:"tags" gorm:"tags"` // TODO: нет необходимости в нём
	Code int    `json:"code" bson:"code" gorm:"code"` // если не 0, то ОШИБКА, читаем лог(Log)
	Log  string `json:"log" bson:"log" gorm:"log"`
}

// type TransData struct --- в apiGetTransaction.go

// type ResponseDeliverTx struct --- в apiGetTransaction.go // TODO: del

// получаем содержимое блока по его ID
func (c *SDK) GetBlock(id int) (BlockResponse, error) {
	url := fmt.Sprintf("%s/api/block/%d", c.MnAddress, id)
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

	for iStep, _ := range data.Result.Transactions {

		if data.Result.Transactions[iStep].Type == TX_SendData {
			data.Result.Transactions[iStep].Data = tx1SendData{
				Coin:  data.Result.Transactions[iStep].DataTx.Coin,
				To:    data.Result.Transactions[iStep].DataTx.To,
				Value: pipStr2bip_f32(data.Result.Transactions[iStep].DataTx.Value),
			}
		} else if data.Result.Transactions[iStep].Type == TX_SellCoinData {
			data.Result.Transactions[iStep].Data = tx2SellCoinData{
				CoinToSell:  data.Result.Transactions[iStep].DataTx.CoinToSell,
				ValueToSell: pipStr2bip_f32(data.Result.Transactions[iStep].DataTx.ValueToSell),
				CoinToBuy:   data.Result.Transactions[iStep].DataTx.CoinToBuy,
			}
			// TODO: Будет-ли в блоке инфа о результате выполнения транзакции Buy/Sell?
			//data.Result.Tags.TxReturn = pipStr2bip_f32(data.Result.Tags.TxReturnTx)
		} else if data.Result.Transactions[iStep].Type == TX_SellAllCoinData {
			data.Result.Transactions[iStep].Data = tx3SellAllCoinData{
				CoinToSell: data.Result.Transactions[iStep].DataTx.CoinToSell,
				CoinToBuy:  data.Result.Transactions[iStep].DataTx.CoinToBuy,
			}
			// TODO: Будет-ли в блоке инфа о результате выполнения транзакции Buy/Sell?
			//data.Result.Tags.TxReturn = pipStr2bip_f32(data.Result.Tags.TxReturnTx)
		} else if data.Result.Transactions[iStep].Type == TX_BuyCoinData {
			data.Result.Transactions[iStep].Data = tx4BuyCoinData{
				CoinToBuy:  data.Result.Transactions[iStep].DataTx.CoinToBuy,
				ValueToBuy: pipStr2bip_f32(data.Result.Transactions[iStep].DataTx.ValueToBuy),
				CoinToSell: data.Result.Transactions[iStep].DataTx.CoinToSell,
			}
			// TODO: Будет-ли в блоке инфа о результате выполнения транзакции Buy/Sell?
			//data.Result.Tags.TxReturn = pipStr2bip_f32(data.Result.Tags.TxReturnTx)
		} else if data.Result.Transactions[iStep].Type == TX_CreateCoinData {
			data.Result.Transactions[iStep].Data = tx5CreateCoinData{
				Name:                 data.Result.Transactions[iStep].DataTx.Name,
				CoinSymbol:           data.Result.Transactions[iStep].DataTx.CoinSymbol,
				InitialAmount:        pipStr2bip_f32(data.Result.Transactions[iStep].DataTx.InitialAmount),
				InitialReserve:       pipStr2bip_f32(data.Result.Transactions[iStep].DataTx.InitialReserve),
				ConstantReserveRatio: data.Result.Transactions[iStep].DataTx.ConstantReserveRatio,
			}
		} else if data.Result.Transactions[iStep].Type == TX_DeclareCandidacyData {
			data.Result.Transactions[iStep].Data = tx6DeclareCandidacyData{
				Address:    data.Result.Transactions[iStep].DataTx.Address,
				PubKey:     data.Result.Transactions[iStep].DataTx.PubKey,
				Commission: data.Result.Transactions[iStep].DataTx.Commission,
				Coin:       data.Result.Transactions[iStep].DataTx.Coin,
				Stake:      pipStr2bip_f32(data.Result.Transactions[iStep].DataTx.Stake),
			}
		} else if data.Result.Transactions[iStep].Type == TX_DelegateDate {
			data.Result.Transactions[iStep].Data = tx7DelegateDate{
				PubKey: data.Result.Transactions[iStep].DataTx.PubKey,
				Coin:   data.Result.Transactions[iStep].DataTx.Coin,
				Stake:  pipStr2bip_f32(data.Result.Transactions[iStep].DataTx.Stake),
			}
		} else if data.Result.Transactions[iStep].Type == TX_UnbondData {
			data.Result.Transactions[iStep].Data = tx8UnbondData{
				PubKey: data.Result.Transactions[iStep].DataTx.PubKey,
				Coin:   data.Result.Transactions[iStep].DataTx.Coin,
				Value:  pipStr2bip_f32(data.Result.Transactions[iStep].DataTx.Value),
			}
		} else if data.Result.Transactions[iStep].Type == TX_RedeemCheckData {
			data.Result.Transactions[iStep].Data = tx9RedeemCheckData{
				RawCheck: data.Result.Transactions[iStep].DataTx.RawCheck,
				Proof:    data.Result.Transactions[iStep].DataTx.Proof,
			}
		} else if data.Result.Transactions[iStep].Type == TX_SetCandidateOnData {
			data.Result.Transactions[iStep].Data = tx10SetCandidateOnData{
				PubKey: data.Result.Transactions[iStep].DataTx.PubKey,
			}
		} else if data.Result.Transactions[iStep].Type == TX_SetCandidateOffData {
			data.Result.Transactions[iStep].Data = tx11SetCandidateOffData{
				PubKey: data.Result.Transactions[iStep].DataTx.PubKey,
			}
		} //else if data.Result.Transactions[iStep].Type == TX_CreateMultisigData {}
	}
	for iStep, _ := range data.Result.Events {
		data.Result.Events[iStep].Value.Amount = pipStr2bip_f32(data.Result.Events[iStep].Value.AmountTx)
	}
	for iStep, _ := range data.Result.Precommits {
		data.Result.Precommits[iStep].ValidatorIndex, err = strconv.Atoi(data.Result.Precommits[iStep].ValidatorIndexTx)
		if err != nil {
			data.Result.Precommits[iStep].ValidatorIndex = 0
		}
		data.Result.Precommits[iStep].Height, err = strconv.Atoi(data.Result.Precommits[iStep].HeightTx)
		if err != nil {
			data.Result.Precommits[iStep].Height = 0
		}
		data.Result.Precommits[iStep].Round, err = strconv.Atoi(data.Result.Precommits[iStep].RoundTx)
		if err != nil {
			data.Result.Precommits[iStep].Round = 0
		}
	}

	return data.Result, nil
}
