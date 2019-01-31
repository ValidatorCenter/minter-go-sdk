package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type node_coininfo struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  CoinInfoResponse
	Error   ErrorStruct
}

type CoinInfoResponse struct {
	Name             string  `json:"name" bson:"name" gorm:"name" db:"name"`
	Symbol           string  `json:"symbol" bson:"symbol" gorm:"symbol" db:"symbol"`
	VolumeTx         string  `json:"volume" bson:"-" gorm:"-" db:"-"`
	Volume           float32 `json:"volume_f32" bson:"volume_f32" gorm:"volume_f32" db:"volume_f32"`
	CRRTx            string  `json:"crr" bson:"-" gorm:"-" db:"-"`
	CRR              int     `json:"crr_i32" bson:"crr_i32" gorm:"crr_i32" db:"crr_i32"`
	ReserveBalanceTx string  `json:"reserve_balance" bson:"-" gorm:"-" db:"-"`
	ReserveBalance   float32 `json:"reserve_balance_f32" bson:"reserve_balance_f32" gorm:"reserve_balance_f32" db:"reserve_balance_f32"`
	//Creator          string  `json:"creator" bson:"creator" gorm:"creator" db:"creator"` // TODO: del, нету
}

// получение доп.данных о монете: volume, reserve_balance
func (c *SDK) GetCoinInfo(coinSmbl string) (CoinInfoResponse, error) {
	url := fmt.Sprintf("%s/coin_info?symbol=%s", c.MnAddress, coinSmbl)
	res, err := http.Get(url)
	if err != nil {
		return CoinInfoResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return CoinInfoResponse{}, err
	}

	var data node_coininfo
	json.Unmarshal(body, &data)

	data.Result.Volume = pipStr2bip_f32(data.Result.VolumeTx)
	data.Result.ReserveBalance = pipStr2bip_f32(data.Result.ReserveBalanceTx)

	data.Result.CRR, err = strconv.Atoi(data.Result.CRRTx)
	if err != nil {
		data.Result.CRR = 0
	}

	return data.Result, nil
}
