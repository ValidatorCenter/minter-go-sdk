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
		data.Result.Transactions[iStep].Height = data.Result.Height
		data.Result.Transactions[iStep].Nonce, err = strconv.Atoi(data.Result.Transactions[iStep].NonceTx)
		if err != nil {
			c.DebugLog("ERROR", "GetBlock-> strconv.Atoi(data.Result.Transactions[iStep].NonceTx)", data.Result.Transactions[iStep].NonceTx)
			return data.Result, err
		}
		data.Result.Transactions[iStep].GasPrice, err = strconv.Atoi(data.Result.Transactions[iStep].GasPriceTx)
		if err != nil {
			c.DebugLog("ERROR", "GetBlock-> strconv.Atoi(data.Result.Transactions[iStep].GasPriceTx)", data.Result.Transactions[iStep].GasPriceTx)
			return data.Result, err
		}
		data.Result.Transactions[iStep].GasUsed, err = strconv.Atoi(data.Result.Transactions[iStep].GasUsedTx)
		if err != nil {
			c.DebugLog("ERROR", "GetBlock-> strconv.Atoi(data.Result.Transactions[iStep].GasUsedTx)", data.Result.Transactions[iStep].GasUsedTx)
			return data.Result, err
		}

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

	return data.Result, nil
}
