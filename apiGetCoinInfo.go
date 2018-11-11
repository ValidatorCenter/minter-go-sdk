package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type node_coininfo struct {
	Code   int
	Result CoinInfoResponse
}

type CoinInfoResponse struct {
	Name           string `json:"name" bson:"name" gorm:"name"`
	Symbol         string `json:"symbol" bson:"symbol" gorm:"symbol"`
	Volume         string `json:"volume" bson:"volume" gorm:"volume"`
	CRR            int    `json:"crr" bson:"crr" gorm:"crr"`
	ReserveBalance string `json:"reserve_balance" bson:"reserve_balance" gorm:"reserve_balance"`
	Creator        string `json:"creator" bson:"creator" gorm:"creator"`
}

// получение доп.данных о монете: volume, reserve_balance
func (c *SDK) GetCoinInfo(coinSmbl string) CoinInfoResponse {
	url := fmt.Sprintf("%s/api/coinInfo/%s", c.MnAddress, coinSmbl)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data node_coininfo
	json.Unmarshal(body, &data)

	return data.Result
}
