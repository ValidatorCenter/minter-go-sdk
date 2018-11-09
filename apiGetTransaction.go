package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"math/big"
	"net/http"
)

//curl -s 'localhost:8841/api/transaction/{hash}'
type node_transaction struct {
	Code   int
	Result TransResponse
}

type TransResponse struct {
	Hash     string      `json:"hash"`
	RawTx    string      `json:"raw_tx"`
	Height   int         `json:"height"`
	Index    int         `json:"index"`
	From     string      `json:"from"`
	Nonce    int         `json:"nonce"`
	GasPrice int         `json:"gas_price"`
	GasCoin  string      `json:"gas_coin"`
	GasUsed  int         `json:"gas_used"`
	Type     int         `json:"type"`
	DataTx   TransData   `json:"data"`
	Data     interface{} `json:"-"`
	Payload  string      `json:"payload"`
	//Tags     TagKeyValue2 `json:"tags"` // TODO: нет необходимости в нём
	Code int    `json:"code"` // (!)не везде
	Log  string `json:"log"`  // (!)не везде
}

// УБРАЛ:
//TxResult ResponseDeliverTx `json:"tx_result"` // TODO: del

/*
tags(send-1):
tx.coin	"MNT"
tx.from	"74996099dfc1aef51761bd6fc52f4e6f6fd75cab"
tx.to	"198208eb4d11d4b389ff262ac52494d920770879"
tx.type	"01"

tx3SellAllCoinData:
tx.coin_to_buy	"CINEMACOIN"
tx.coin_to_sell	"MNT"
tx.from	"3c57a889ec01714f26477f3758ee3b5c08bcabd3"
tx.return	"511368399233653912755"
tx.sell_amount	"968023591980876730077"
tx.type	"02"

tx5CreateCoinData:
tx.coin	"DOGECOIN"
tx.from	"89f5395a03847826d6b48bb02dbde64376945a20"
tx.type	"05"
*/

/*
TODO: del
type ResponseDeliverTx struct {
	GasWanted int64        `json:"gas_wanted"`
	GasUsed   int64        `json:"gas_used"`
	Tags      TagKeyValue2 `json:"tags" bson:"tags"`
}*/

type TagKeyValue2 struct {
	TxCoinToBuy  string `json:"tx.coin_to_buy"`
	TxCoinToSell string `json:"tx.coin_to_sell"`
	TxFrom       string `json:"tx.from"`
	TxReturn     string `json:"tx.return"`
	TxSellAmount string `json:"tx.sell_amount"`
	//tx.type	"\u0002"
}

type tx1SendData struct {
	Coin  string `json:"coin"`
	To    string `json:"to"`
	Value string `json:"value"`
}

type tx2SellCoinData struct {
	CoinToSell  string `json:"coin_to_sell"`
	ValueToSell string `json:"value_to_sell"`
	CoinToBuy   string `json:"coin_to_buy"`
}

type tx3SellAllCoinData struct {
	CoinToSell string `json:"coin_to_sell"`
	CoinToBuy  string `json:"coin_to_buy"`
}

type tx4BuyCoinData struct {
	CoinToBuy  string `json:"coin_to_buy"`
	ValueToBuy string `json:"value_to_buy"`
	CoinToSell string `json:"coin_to_sell"`
}

type tx5CreateCoinData struct {
	Name                 string `json:"name"`
	CoinSymbol           string `json:"coin_symbol"`
	InitialAmount        string `json:"initial_amount"`
	InitialReserve       string `json:"initial_reserve"`
	ConstantReserveRatio int    `json:"constant_reserve_ratio"`
}

type tx6DeclareCandidacyData struct {
	Address    string `json:"address"`
	PubKey     string `json:"pub_key"`
	Commission int    `json:"commission"`
	Coin       string `json:"coin"`
	Stake      string `json:"stake"`
}

type tx7DelegateDate struct {
	PubKey string `json:"pub_key"`
	Coin   string `json:"coin"`
	Stake  string `json:"stake"`
}

type tx8UnbondData struct {
	PubKey string `json:"pub_key"`
	Coin   string `json:"coin"`
	Value  string `json:"value"`
}

type tx9RedeemCheckData struct {
	RawCheck string `json:"raw_check"`
	Proof    string `json:"proof"`
}

type tx10SetCandidateOnData struct {
	PubKey string `json:"pub_key"`
}

type tx11SetCandidateOffData struct {
	PubKey string `json:"pub_key"`
}

type tx12CreateMultisigData struct {
	/*Threshold uint
	Weights   []uint
	Addresses [][20]byte*/
}

type TransData struct {
	//=== type1 - TYPE_SEND
	To string `json:"to"`
	//Coin   string `json:"coin"`
	//Value  string `json:"value"`
	//=== type2 - TYPE_SELL_COIN
	ValueToSell string `json:"value_to_sell"`
	//CoinToSell string `json:"coin_to_sell"`
	//CoinToBuy  string `json:"coin_to_buy"`
	//=== type3 - TYPE_SELL_ALL_COIN
	//CoinToSell string `json:"coin_to_sell"`
	//CoinToBuy  string `json:"coin_to_buy"`
	//=== type4 - TYPE_BUY_COIN
	CoinToBuy  string `json:"coin_to_buy"`
	ValueToBuy string `json:"value_to_buy"`
	CoinToSell string `json:"coin_to_sell"`
	//=== type5 - TYPE_CREATE_COIN
	Name                 string `json:"name"`                   // название монеты
	CoinSymbol           string `json:"coin_symbol"`            // символ монеты
	InitialAmount        string `json:"initial_amount"`         //  Amount of coins to issue. Issued coins will be available to sender account.
	InitialReserve       string `json:"initial_reserve"`        // Initial reserve in base coin.
	ConstantReserveRatio int    `json:"constant_reserve_ratio"` // uint, should be from 10 to 100 (в %).
	//=== type6 - TYPE_DECLARE_CANDIDACY
	Address    string `json:"address"`
	Commission int    `json:"commission"`
	//Stake	string `json:"stake"`
	//PubKey  string `json:"pub_key"`
	//Coin    string `json:"coin"`
	//=== type7 - TYPE_DELEGATE
	Stake string `json:"stake"`
	//PubKey  string `json:"pub_key"`
	//Coin    string `json:"coin"`
	//=== type8 - TYPE_UNBOUND
	PubKey string `json:"pub_key"`
	Coin   string `json:"coin"`
	Value  string `json:"value"`
	//=== type9 - TYPE_REDEEM_CHECK
	RawCheck string `json:"raw_check"`
	Proof    string `json:"proof"`
	//=== type10 - TYPE_SET_CANDIDATE_ONLINE
	//=== type11 - TYPE_SET_CANDIDATE_OFFLINE
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
	fmt.Println(string(body))

	if data.Result.Type == 1 {
		data.Result.Data = tx1SendData{
			Coin:  data.Result.DataTx.Coin,
			To:    data.Result.DataTx.To,
			Value: data.Result.DataTx.Value,
		}
	} else if data.Result.Type == 2 {
		data.Result.Data = tx2SellCoinData{
			CoinToSell:  data.Result.DataTx.CoinToSell,
			ValueToSell: data.Result.DataTx.ValueToSell,
			CoinToBuy:   data.Result.DataTx.CoinToBuy,
		}
	} else if data.Result.Type == 3 {
		data.Result.Data = tx3SellAllCoinData{
			CoinToSell: data.Result.DataTx.CoinToSell,
			CoinToBuy:  data.Result.DataTx.CoinToBuy,
		}
	} else if data.Result.Type == 4 {
		data.Result.Data = tx4BuyCoinData{
			CoinToBuy:  data.Result.DataTx.CoinToBuy,
			ValueToBuy: data.Result.DataTx.ValueToBuy,
			CoinToSell: data.Result.DataTx.CoinToSell,
		}
	} else if data.Result.Type == 5 {
		data.Result.Data = tx5CreateCoinData{
			Name:                 data.Result.DataTx.Name,
			CoinSymbol:           data.Result.DataTx.CoinSymbol,
			InitialAmount:        data.Result.DataTx.InitialAmount,
			InitialReserve:       data.Result.DataTx.InitialReserve,
			ConstantReserveRatio: data.Result.DataTx.ConstantReserveRatio,
		}
	} else if data.Result.Type == 6 {
		data.Result.Data = tx6DeclareCandidacyData{
			Address:    data.Result.DataTx.Address,
			PubKey:     data.Result.DataTx.PubKey,
			Commission: data.Result.DataTx.Commission,
			Coin:       data.Result.DataTx.Coin,
			Stake:      data.Result.DataTx.Stake,
		}
	} else if data.Result.Type == 7 {
		data.Result.Data = tx7DelegateDate{
			PubKey: data.Result.DataTx.PubKey,
			Coin:   data.Result.DataTx.Coin,
			Stake:  data.Result.DataTx.Stake,
		}
	} else if data.Result.Type == 8 {
		data.Result.Data = tx8UnbondData{
			PubKey: data.Result.DataTx.PubKey,
			Coin:   data.Result.DataTx.Coin,
			Value:  data.Result.DataTx.Value,
		}
	} else if data.Result.Type == 9 {
		data.Result.Data = tx9RedeemCheckData{
			RawCheck: data.Result.DataTx.RawCheck,
			Proof:    data.Result.DataTx.Proof,
		}
	} else if data.Result.Type == 10 {
		data.Result.Data = tx10SetCandidateOnData{
			PubKey: data.Result.DataTx.PubKey,
		}
	} else if data.Result.Type == 11 {
		data.Result.Data = tx11SetCandidateOffData{
			PubKey: data.Result.DataTx.PubKey,
		}
	} //else if data.Result.Type == 12 {}

	return data.Result
}
