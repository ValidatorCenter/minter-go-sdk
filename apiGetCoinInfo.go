package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type node_coininfo struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  CoinInfoResponse
}

type CoinInfoResponse struct {
	Name             string  `json:"name" bson:"name" gorm:"name"`
	Symbol           string  `json:"symbol" bson:"symbol" gorm:"symbol"`
	VolumeTx         string  `json:"volume" bson:"-" gorm:"-"`
	Volume           float32 `json:"volume_f32" bson:"volume_f32" gorm:"volume_f32"`
	CRR              int     `json:"crr" bson:"crr" gorm:"crr"`
	ReserveBalanceTx string  `json:"reserve_balance" bson:"-" gorm:"-"`
	ReserveBalance   float32 `json:"reserve_balance_f32" bson:"reserve_balance_f32" gorm:"reserve_balance_f32"`
	//Creator          string  `json:"creator" bson:"creator" gorm:"creator"` // TODO: del, нету
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

	return data.Result, nil
}
