package mintersdk

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type node_transaction struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  TransResponse
}

type TransResponse struct {
	Hash     string       `json:"hash" bson:"hash" gorm:"hash"`
	RawTx    string       `json:"raw_tx" bson:"raw_tx" gorm:"raw_tx"`
	Height   int          `json:"height" bson:"height" gorm:"height"`
	Index    int          `json:"index" bson:"index" gorm:"index"`
	From     string       `json:"from" bson:"from" gorm:"from"`
	Nonce    int          `json:"nonce" bson:"nonce" gorm:"nonce"`
	GasPrice int          `json:"gas_price" bson:"gas_price" gorm:"gas_price"`
	GasCoin  string       `json:"gas_coin" bson:"gas_coin" gorm:"gas_coin"`
	GasUsed  int          `json:"gas_used" bson:"gas_used" gorm:"gas_used"`
	Type     int          `json:"type" bson:"type" gorm:"type"`
	DataTx   TransData    `json:"data" bson:"-" gorm:"-"`
	Data     interface{}  `json:"-" bson:"data" gorm:"data"`
	Payload  string       `json:"payload" bson:"payload" gorm:"payload"`
	Tags     tagKeyValue2 `json:"tags" bson:"tags" gorm:"tags"` // TODO: нет необходимости в нём, пока из Покупки/Продажи результат обмена tx.return не вынесут на уровень выше
	Code     int          `json:"code" bson:"code" gorm:"code"` // если не 0, то ОШИБКА, читаем лог(Log)
	Log      string       `json:"log" bson:"log" gorm:"log"`
}

// УБРАЛ:
//TxResult ResponseDeliverTx `json:"tx_result"` // TODO: del

/*
TODO: del
type ResponseDeliverTx struct {
	GasWanted int64        `json:"gas_wanted"`
	GasUsed   int64        `json:"gas_used"`
	Tags      TagKeyValue2 `json:"tags" bson:"tags"`
}*/

type tagKeyValue2 struct {
	TxCoinToBuy    string  `json:"tx.coin_to_buy" bson:"tx_coin_to_buy" gorm:"tx_coin_to_buy"`
	TxCoinToSell   string  `json:"tx.coin_to_sell" bson:"tx_coin_to_sell" gorm:"tx_coin_to_sell"`
	TxFrom         string  `json:"tx.from" bson:"tx_from" gorm:"tx_from"`
	TxReturnTx     string  `json:"tx.return" bson:"-" gorm:"-"`
	TxReturn       float32 `json:"tx.return_f32" bson:"tx_return_f32" gorm:"tx_return_f32"`
	TxSellAmountTx string  `json:"tx.sell_amount" bson:"-" gorm:"-"`
	TxSellAmount   float32 `json:"tx.sell_amount_f32" bson:"tx_sell_amount_f32" gorm:"tx_sell_amount_f32"`
	//tx.type	"\u0002"
}

type tx1SendData struct {
	Coin string `json:"coin" bson:"coin" gorm:"coin"`
	To   string `json:"to" bson:"to" gorm:"to"`
	//ValueTx string  `json:"value" bson:"-" gorm:"-"`
	Value float32 `json:"value_f32" bson:"value_f32" gorm:"value_f32"`
}

type tx2SellCoinData struct {
	CoinToSell string `json:"coin_to_sell" bson:"coin_to_sell" gorm:"coin_to_sell"`
	CoinToBuy  string `json:"coin_to_buy" bson:"coin_to_buy" gorm:"coin_to_buy"`
	//ValueToSellTx string  `json:"value_to_sell" bson:"-" gorm:"-"`
	ValueToSell float32 `json:"value_to_sell_f32" bson:"value_to_sell_f32" gorm:"value_to_sell_f32"`
}

type tx3SellAllCoinData struct {
	CoinToSell string `json:"coin_to_sell" bson:"coin_to_sell" gorm:"coin_to_sell"`
	CoinToBuy  string `json:"coin_to_buy" bson:"coin_to_buy" gorm:"coin_to_buy"`
}

type tx4BuyCoinData struct {
	CoinToBuy  string `json:"coin_to_buy" bson:"coin_to_buy" gorm:"coin_to_buy"`
	CoinToSell string `json:"coin_to_sell" bson:"coin_to_sell" gorm:"coin_to_sell"`
	//ValueToBuyTx string  `json:"value_to_buy" bson:"-" gorm:"-"`
	ValueToBuy float32 `json:"value_to_buy_f32" bson:"value_to_buy_f32" gorm:"value_to_buy_f32"`
}

type tx5CreateCoinData struct {
	Name       string `json:"name" bson:"name" gorm:"name"`
	CoinSymbol string `json:"coin_symbol" bson:"coin_symbol" gorm:"coin_symbol"`
	//InitialAmountTx      string  `json:"initial_amount" bson:"-" gorm:"-"`
	//InitialReserveTx     string  `json:"initial_reserve" bson:"-" gorm:"-"`
	ConstantReserveRatio int     `json:"constant_reserve_ratio" bson:"constant_reserve_ratio" gorm:"constant_reserve_ratio"`
	InitialAmount        float32 `json:"initial_amount_f32" bson:"initial_amount_f32" gorm:"initial_amount_f32"`
	InitialReserve       float32 `json:"initial_reserve_f32" bson:"initial_reserve_f32" gorm:"initial_reserve_f32"`
}

type tx6DeclareCandidacyData struct {
	Address    string `json:"address" bson:"address" gorm:"address"`
	PubKey     string `json:"pub_key" bson:"pub_key" gorm:"pub_key"`
	Commission int    `json:"commission" bson:"commission" gorm:"commission"`
	Coin       string `json:"coin" bson:"coin" gorm:"coin"`
	//StakeTx    string  `json:"stake" bson:"-" gorm:"-"`
	Stake float32 `json:"stake_f32" bson:"stake_f32" gorm:"stake_f32"`
}

type tx7DelegateDate struct {
	PubKey string `json:"pub_key" bson:"pub_key" gorm:"pub_key"`
	Coin   string `json:"coin" bson:"coin" gorm:"coin"`
	//StakeTx string  `json:"stake" bson:"-" gorm:"-"`
	Stake float32 `json:"stake_f32" bson:"stake_f32" gorm:"stake_f32"`
}

type tx8UnbondData struct {
	PubKey string `json:"pub_key" bson:"pub_key" gorm:"pub_key"`
	Coin   string `json:"coin" bson:"coin" gorm:"coin"`
	//ValueTx string  `json:"value" bson:"-" gorm:"-"`
	Value float32 `json:"value_f32" bson:"value_f32" gorm:"value_f32"`
}

type tx9RedeemCheckData struct {
	RawCheck string `json:"raw_check" bson:"raw_check" gorm:"raw_check"`
	Proof    string `json:"proof" bson:"proof" gorm:"proof"`
}

type tx10SetCandidateOnData struct {
	PubKey string `json:"pub_key" bson:"pub_key" gorm:"pub_key"`
}

type tx11SetCandidateOffData struct {
	PubKey string `json:"pub_key" bson:"pub_key" gorm:"pub_key"`
}

type tx12CreateMultisigData struct {
	/**
	Threshold uint
	Weights   []uint
	Addresses [][20]byte
	**/
}

// Не заносится в БД
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
func (c *SDK) GetTransaction(hash string) (TransResponse, error) {
	url := fmt.Sprintf("%s/transaction?hash=%s", c.MnAddress, hash)
	res, err := http.Get(url)
	if err != nil {
		return TransResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return TransResponse{}, err
	}

	var data node_transaction
	json.Unmarshal(body, &data)
	fmt.Println(string(body))

	if data.Result.Type == TX_SendData {
		data.Result.Data = tx1SendData{
			Coin:  data.Result.DataTx.Coin,
			To:    data.Result.DataTx.To,
			Value: pipStr2bip_f32(data.Result.DataTx.Value),
		}
	} else if data.Result.Type == TX_SellCoinData {
		data.Result.Data = tx2SellCoinData{
			CoinToSell:  data.Result.DataTx.CoinToSell,
			ValueToSell: pipStr2bip_f32(data.Result.DataTx.ValueToSell),
			CoinToBuy:   data.Result.DataTx.CoinToBuy,
		}
		data.Result.Tags.TxReturn = pipStr2bip_f32(data.Result.Tags.TxReturnTx)
		data.Result.Tags.TxSellAmount = pipStr2bip_f32(data.Result.Tags.TxSellAmountTx)
	} else if data.Result.Type == TX_SellAllCoinData {
		data.Result.Data = tx3SellAllCoinData{
			CoinToSell: data.Result.DataTx.CoinToSell,
			CoinToBuy:  data.Result.DataTx.CoinToBuy,
		}
		data.Result.Tags.TxReturn = pipStr2bip_f32(data.Result.Tags.TxReturnTx)
		data.Result.Tags.TxSellAmount = pipStr2bip_f32(data.Result.Tags.TxSellAmountTx)
	} else if data.Result.Type == TX_BuyCoinData {
		data.Result.Data = tx4BuyCoinData{
			CoinToBuy:  data.Result.DataTx.CoinToBuy,
			ValueToBuy: pipStr2bip_f32(data.Result.DataTx.ValueToBuy),
			CoinToSell: data.Result.DataTx.CoinToSell,
		}
		data.Result.Tags.TxReturn = pipStr2bip_f32(data.Result.Tags.TxReturnTx)
		data.Result.Tags.TxSellAmount = pipStr2bip_f32(data.Result.Tags.TxSellAmountTx)
	} else if data.Result.Type == TX_CreateCoinData {
		data.Result.Data = tx5CreateCoinData{
			Name:                 data.Result.DataTx.Name,
			CoinSymbol:           data.Result.DataTx.CoinSymbol,
			InitialAmount:        pipStr2bip_f32(data.Result.DataTx.InitialAmount),
			InitialReserve:       pipStr2bip_f32(data.Result.DataTx.InitialReserve),
			ConstantReserveRatio: data.Result.DataTx.ConstantReserveRatio,
		}
	} else if data.Result.Type == TX_DeclareCandidacyData {
		data.Result.Data = tx6DeclareCandidacyData{
			Address:    data.Result.DataTx.Address,
			PubKey:     data.Result.DataTx.PubKey,
			Commission: data.Result.DataTx.Commission,
			Coin:       data.Result.DataTx.Coin,
			Stake:      pipStr2bip_f32(data.Result.DataTx.Stake),
		}
	} else if data.Result.Type == TX_DelegateDate {
		data.Result.Data = tx7DelegateDate{
			PubKey: data.Result.DataTx.PubKey,
			Coin:   data.Result.DataTx.Coin,
			Stake:  pipStr2bip_f32(data.Result.DataTx.Stake),
		}
	} else if data.Result.Type == TX_UnbondData {
		data.Result.Data = tx8UnbondData{
			PubKey: data.Result.DataTx.PubKey,
			Coin:   data.Result.DataTx.Coin,
			Value:  pipStr2bip_f32(data.Result.DataTx.Value),
		}
	} else if data.Result.Type == TX_RedeemCheckData {
		data.Result.Data = tx9RedeemCheckData{
			RawCheck: data.Result.DataTx.RawCheck,
			Proof:    data.Result.DataTx.Proof,
		}
	} else if data.Result.Type == TX_SetCandidateOnData {
		data.Result.Data = tx10SetCandidateOnData{
			PubKey: data.Result.DataTx.PubKey,
		}
	} else if data.Result.Type == TX_SetCandidateOffData {
		data.Result.Data = tx11SetCandidateOffData{
			PubKey: data.Result.DataTx.PubKey,
		}
	} else if data.Result.Type == TX_CreateMultisigData {
		// TODO: реализовать
	}

	// Расшифровываем сообщение
	if data.Result.Payload != "" {
		// комментарий, расшифровать base64
		sDec, _ := b64.StdEncoding.DecodeString(data.Result.Payload)
		data.Result.Payload = string(sDec)
	}

	return data.Result, nil
}
