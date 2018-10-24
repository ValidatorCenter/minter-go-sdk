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
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	Volume         string `json:"volume"`
	CRR            int    `json:"crr"`
	ReserveBalance string `json:"reserve_balance"`
	Creator        string `json:"creator"`
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
