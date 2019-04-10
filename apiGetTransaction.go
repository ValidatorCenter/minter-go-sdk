package mintersdk

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type node_transaction struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  TransResponse
	Error   ErrorStruct
}

type TransResponse struct {
	Hash        string       `json:"hash" bson:"hash" gorm:"hash" db:"hash"`
	RawTx       string       `json:"raw_tx" bson:"raw_tx" gorm:"raw_tx" db:"raw_tx"`
	HeightTx    string       `json:"height" bson:"-" gorm:"-" db:"-"`
	Height      int          `json:"height_i32" bson:"height_i32" gorm:"height_i32" db:"height_i32"` //(!) В блоке у транзакции нет HEIGHT блока
	Index       int          `json:"index" bson:"index" gorm:"index" db:"index_i32"`
	From        string       `json:"from" bson:"from" gorm:"from" db:"from"`
	NonceTx     string       `json:"nonce" bson:"-" gorm:"-" db:"-"`
	Nonce       int          `json:"nonce_i32" bson:"nonce_i32" gorm:"nonce_i32" db:"nonce_i32"`
	GasPriceTx  string       `json:"gas_price" bson:"-" gorm:"-" db:"-"`
	GasPrice    int          `json:"gas_price_i32" bson:"gas_price_i32" gorm:"gas_price_i32" db:"gas_price_i32"`
	GasCoin     string       `json:"gas_coin" bson:"gas_coin" gorm:"gas_coin" db:"gas_coin"`
	GasUsedTx   string       `json:"gas_used" bson:"-" gorm:"-" db:"-"`
	GasUsed     int          `json:"gas_used_i32" bson:"gas_used_i32" gorm:"gas_used_i32" db:"gas_used_i32"`
	Type        int          `json:"type" bson:"type" gorm:"type" db:"type"`
	DataTx      TransData    `json:"data" bson:"-" gorm:"-" db:"-"`
	Data        interface{}  `json:"-" bson:"data" gorm:"data" db:"data"`
	Payload     string       `json:"payload" bson:"payload" gorm:"payload" db:"payload"`
	Tags        tagKeyValue2 `json:"tags" bson:"tags" gorm:"tags" db:"tags"` // TODO: нет необходимости в нём, пока из Покупки/Продажи результат обмена tx.return не вынесут на уровень выше
	Code        int          `json:"code" bson:"code" gorm:"code" db:"code"` // если не 0, то ОШИБКА, читаем лог(Log)
	Log         string       `json:"log" bson:"log" gorm:"log" db:"log"`
	ServiceData []byte       `json:"service_data" bson:"service_data" gorm:"service_data" db:"service_data"` //TODO: ?
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
	TxCoinToBuy    string  `json:"tx.coin_to_buy" bson:"tx_coin_to_buy" gorm:"tx_coin_to_buy" db:"tx_coin_to_buy"`
	TxCoinToSell   string  `json:"tx.coin_to_sell" bson:"tx_coin_to_sell" gorm:"tx_coin_to_sell" db:"tx_coin_to_sell"`
	TxFrom         string  `json:"tx.from" bson:"tx_from" gorm:"tx_from" db:"tx_from"`
	TxReturnTx     string  `json:"tx.return" bson:"-" gorm:"-" db:"-"`
	TxReturn       float32 `json:"tx.return_f32" bson:"tx_return_f32" gorm:"tx_return_f32" db:"tx_return_f32"`
	TxSellAmountTx string  `json:"tx.sell_amount" bson:"-" gorm:"-" db:"-"`
	TxSellAmount   float32 `json:"tx.sell_amount_f32" bson:"tx_sell_amount_f32" gorm:"tx_sell_amount_f32" db:"tx_sell_amount_f32"`
	//tx.type	"\u0002"
}

type tx1SendData struct {
	Coin string `json:"coin" bson:"coin" gorm:"coin" db:"coin"`
	To   string `json:"to" bson:"to" gorm:"to" db:"to"`
	//ValueTx string  `json:"value" bson:"-" gorm:"-" db:"-"`
	Value float32 `json:"value_f32" bson:"value_f32" gorm:"value_f32" db:"value_f32"`
}

type tx2SellCoinData struct {
	CoinToSell string `json:"coin_to_sell" bson:"coin_to_sell" gorm:"coin_to_sell" db:"coin_to_sell"`
	CoinToBuy  string `json:"coin_to_buy" bson:"coin_to_buy" gorm:"coin_to_buy" db:"coin_to_buy"`
	//ValueToSellTx string  `json:"value_to_sell" bson:"-" gorm:"-" db:"-"`
	ValueToSell float32 `json:"value_to_sell_f32" bson:"value_to_sell_f32" gorm:"value_to_sell_f32" db:"value_to_sell_f32"`
}

type tx3SellAllCoinData struct {
	CoinToSell string `json:"coin_to_sell" bson:"coin_to_sell" gorm:"coin_to_sell" db:"coin_to_sell"`
	CoinToBuy  string `json:"coin_to_buy" bson:"coin_to_buy" gorm:"coin_to_buy" db:"coin_to_buy"`
}

type tx4BuyCoinData struct {
	CoinToBuy  string `json:"coin_to_buy" bson:"coin_to_buy" gorm:"coin_to_buy" db:"coin_to_buy"`
	CoinToSell string `json:"coin_to_sell" bson:"coin_to_sell" gorm:"coin_to_sell" db:"coin_to_sell"`
	//ValueToBuyTx string  `json:"value_to_buy" bson:"-" gorm:"-" db:"-"`
	ValueToBuy float32 `json:"value_to_buy_f32" bson:"value_to_buy_f32" gorm:"value_to_buy_f32" db:"value_to_buy_f32"`
}

type tx5CreateCoinData struct {
	Name       string `json:"name" bson:"name" gorm:"name" db:"name"`
	CoinSymbol string `json:"symbol" bson:"symbol" gorm:"symbol" db:"symbol"`
	//InitialAmountTx      string  `json:"initial_amount" bson:"-" gorm:"-" db:"-"`
	//InitialReserveTx     string  `json:"initial_reserve" bson:"-" gorm:"-" db:"-"`
	ConstantReserveRatio int     `json:"constant_reserve_ratio" bson:"constant_reserve_ratio" gorm:"constant_reserve_ratio" db:"constant_reserve_ratio"`
	InitialAmount        float32 `json:"initial_amount_f32" bson:"initial_amount_f32" gorm:"initial_amount_f32" db:"initial_amount_f32"`
	InitialReserve       float32 `json:"initial_reserve_f32" bson:"initial_reserve_f32" gorm:"initial_reserve_f32" db:"initial_reserve_f32"`
}

type tx6DeclareCandidacyData struct {
	Address    string `json:"address" bson:"address" gorm:"address" db:"address"`
	PubKey     string `json:"pub_key" bson:"pub_key" gorm:"pub_key" db:"pub_key"`
	Commission int    `json:"commission" bson:"commission" gorm:"commission" db:"commission"`
	Coin       string `json:"coin" bson:"coin" gorm:"coin" db:"coin"`
	//StakeTx    string  `json:"stake" bson:"-" gorm:"-" db:"-"`
	Stake float32 `json:"stake_f32" bson:"stake_f32" gorm:"stake_f32" db:"stake_f32"`
}

type tx7DelegateDate struct {
	PubKey string `json:"pub_key" bson:"pub_key" gorm:"pub_key" db:"pub_key"`
	Coin   string `json:"coin" bson:"coin" gorm:"coin" db:"coin"`
	//StakeTx string  `json:"stake" bson:"-" gorm:"-" db:"-"`
	Stake float32 `json:"stake_f32" bson:"stake_f32" gorm:"stake_f32" db:"stake_f32"`
}

type tx8UnbondData struct {
	PubKey string `json:"pub_key" bson:"pub_key" gorm:"pub_key" db:"pub_key"`
	Coin   string `json:"coin" bson:"coin" gorm:"coin" db:"coin"`
	//ValueTx string  `json:"value" bson:"-" gorm:"-" db:"-"`
	Value float32 `json:"value_f32" bson:"value_f32" gorm:"value_f32" db:"value_f32"`
}

type tx9RedeemCheckData struct {
	RawCheck string `json:"raw_check" bson:"raw_check" gorm:"raw_check" db:"raw_check"`
	Proof    string `json:"proof" bson:"proof" gorm:"proof" db:"proof"`
}

type tx10SetCandidateOnData struct {
	PubKey string `json:"pub_key" bson:"pub_key" gorm:"pub_key" db:"pub_key"`
}

type tx11SetCandidateOffData struct {
	PubKey string `json:"pub_key" bson:"pub_key" gorm:"pub_key" db:"pub_key"`
}

type tx12CreateMultisigData struct {
	/**
	Threshold uint
	Weights   []uint
	Addresses [][20]byte
	**/
}

type tx13MultisendData struct {
	List []tx1SendData `json:"list" bson:"list" gorm:"list" db:"list"`
}

// Не заносится в БД
type SendOneData struct {
	To    string `json:"to"`
	Coin  string `json:"coin"`
	Value string `json:"value"`
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
	CoinSymbol           string `json:"symbol"`                 // символ монеты
	InitialAmount        string `json:"initial_amount"`         //  Amount of coins to issue. Issued coins will be available to sender account.
	InitialReserve       string `json:"initial_reserve"`        // Initial reserve in base coin.
	ConstantReserveRatio string `json:"constant_reserve_ratio"` // should be from 10 to 100 (в %).
	//=== type6 - TYPE_DECLARE_CANDIDACY
	Address    string `json:"address"`
	Commission int    `json:"commission,string"`
	//Stake	string `json:"stake"`
	//PubKey  string `json:"pub_key"`
	//Coin    string `json:"coin"`
	//=== type7 - TYPE_DELEGATE
	Stake string `json:"stake"` // УДАЛИТЬ: с 0.15.* стало Value  string `json:"value"`
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
	//=== type13 - TYPE_MULTISEND
	List []SendOneData `json:"list"`
}

// обработка данных транзакции
func manipulationTransaction(c *SDK, tr *TransResponse) error {
	var err error
	tr.Height, err = strconv.Atoi(tr.HeightTx)
	if err != nil {
		c.DebugLog("ERROR", "GetTransaction-> strconv.Atoi(tr.HeightTx)", tr.HeightTx)
		return err
	}
	tr.Nonce, err = strconv.Atoi(tr.NonceTx)
	if err != nil {
		c.DebugLog("ERROR", "GetTransaction-> strconv.Atoi(tr.NonceTx)", tr.NonceTx)
		return err
	}
	tr.GasPrice, err = strconv.Atoi(tr.GasPriceTx)
	if err != nil {
		c.DebugLog("ERROR", "GetTransaction-> strconv.Atoi(tr.GasPriceTx)", tr.GasPriceTx)
		return err
	}
	// TODO: в minter 0.9.x нет
	/*
		tr.GasUsed, err = strconv.Atoi(tr.GasUsedTx)
		if err != nil {
			c.DebugLog("ERROR", "GetTransaction-> strconv.Atoi(tr.GasUsedTx)", tr.GasUsedTx)
			return err
		}*/

	if tr.Type == TX_SendData {
		tr.Data = tx1SendData{
			Coin:  tr.DataTx.Coin,
			To:    tr.DataTx.To,
			Value: pipStr2bip_f32(tr.DataTx.Value),
		}
	} else if tr.Type == TX_SellCoinData {
		tr.Data = tx2SellCoinData{
			CoinToSell:  tr.DataTx.CoinToSell,
			ValueToSell: pipStr2bip_f32(tr.DataTx.ValueToSell),
			CoinToBuy:   tr.DataTx.CoinToBuy,
		}
		tr.Tags.TxReturn = pipStr2bip_f32(tr.Tags.TxReturnTx)
		tr.Tags.TxSellAmount = pipStr2bip_f32(tr.Tags.TxSellAmountTx)
	} else if tr.Type == TX_SellAllCoinData {
		tr.Data = tx3SellAllCoinData{
			CoinToSell: tr.DataTx.CoinToSell,
			CoinToBuy:  tr.DataTx.CoinToBuy,
		}
		tr.Tags.TxReturn = pipStr2bip_f32(tr.Tags.TxReturnTx)
		tr.Tags.TxSellAmount = pipStr2bip_f32(tr.Tags.TxSellAmountTx)
	} else if tr.Type == TX_BuyCoinData {
		tr.Data = tx4BuyCoinData{
			CoinToBuy:  tr.DataTx.CoinToBuy,
			ValueToBuy: pipStr2bip_f32(tr.DataTx.ValueToBuy),
			CoinToSell: tr.DataTx.CoinToSell,
		}
		tr.Tags.TxReturn = pipStr2bip_f32(tr.Tags.TxReturnTx)
		tr.Tags.TxSellAmount = pipStr2bip_f32(tr.Tags.TxSellAmountTx)
	} else if tr.Type == TX_CreateCoinData {
		crrInt, err := strconv.Atoi(tr.DataTx.ConstantReserveRatio)
		if err != nil {
			c.DebugLog("ERROR", "GetTransaction-> strconv.Atoi(tr.DataTx.ConstantReserveRatio)", tr.DataTx.ConstantReserveRatio)
			crrInt = 0
		}
		tr.Data = tx5CreateCoinData{
			Name:                 tr.DataTx.Name,
			CoinSymbol:           tr.DataTx.CoinSymbol,
			InitialAmount:        pipStr2bip_f32(tr.DataTx.InitialAmount),
			InitialReserve:       pipStr2bip_f32(tr.DataTx.InitialReserve),
			ConstantReserveRatio: crrInt,
		}
	} else if tr.Type == TX_DeclareCandidacyData {
		tr.Data = tx6DeclareCandidacyData{
			Address:    tr.DataTx.Address,
			PubKey:     tr.DataTx.PubKey,
			Commission: tr.DataTx.Commission,
			Coin:       tr.DataTx.Coin,
			Stake:      pipStr2bip_f32(tr.DataTx.Stake),
		}
	} else if tr.Type == TX_DelegateDate {
		tr.Data = tx7DelegateDate{
			PubKey: tr.DataTx.PubKey,
			Coin:   tr.DataTx.Coin,
			Stake:  pipStr2bip_f32(tr.DataTx.Value),
		}
	} else if tr.Type == TX_UnbondData {
		tr.Data = tx8UnbondData{
			PubKey: tr.DataTx.PubKey,
			Coin:   tr.DataTx.Coin,
			Value:  pipStr2bip_f32(tr.DataTx.Value),
		}
	} else if tr.Type == TX_RedeemCheckData {
		tr.Data = tx9RedeemCheckData{
			RawCheck: tr.DataTx.RawCheck,
			Proof:    tr.DataTx.Proof,
		}
	} else if tr.Type == TX_SetCandidateOnData {
		tr.Data = tx10SetCandidateOnData{
			PubKey: tr.DataTx.PubKey,
		}
	} else if tr.Type == TX_SetCandidateOffData {
		tr.Data = tx11SetCandidateOffData{
			PubKey: tr.DataTx.PubKey,
		}
	} else if tr.Type == TX_CreateMultisigData {
		// TODO: реализовать
	} else if tr.Type == TX_MultisendData {
		tmpTx13 := tx13MultisendData{}
		for _, itm := range tr.DataTx.List {
			tmpTx13.List = append(tmpTx13.List, tx1SendData{
				Coin:  itm.Coin,
				To:    itm.To,
				Value: pipStr2bip_f32(itm.Value),
			})
		}
		tr.Data = tmpTx13
	} else if tr.Type == TX_EditCandidateData {
		// TODO: реализовать
	}

	// Расшифровываем сообщение
	if tr.Payload != "" {
		// комментарий, расшифровать base64
		sDec, _ := b64.StdEncoding.DecodeString(tr.Payload)
		tr.Payload = string(sDec)
	}
	return nil
}

// получаем содержимое транзакции по её хэшу
func (c *SDK) GetTransaction(hash string) (TransResponse, error) {
	// 0x.. или Mt..
	prefixTx := ""
	if hash[0:2] == "Mt" || hash[0:2] == "0x" {
	} else {
		prefixTx = "0x"
	}
	url := fmt.Sprintf("%s/transaction?hash=%s%s", c.MnAddress, prefixTx, hash)
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

	if data.Error.Code != 0 {
		err = errors.New(fmt.Sprint(data.Error.Code, " - ", data.Error.Message))
		return TransResponse{}, err
	}

	err = manipulationTransaction(c, &data.Result)
	if err != nil {
		return TransResponse{}, err
	}

	return data.Result, nil
}
