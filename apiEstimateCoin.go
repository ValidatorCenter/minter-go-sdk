package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type node_estimate struct {
	Code   int
	Result EstimateResponse
}

// для запроса о стоимости монет
type EstimateResponse struct {
	WillPay    string `json:"will_pay"`
	WillGet    string `json:"will_get"`
	Commission string `json:"commission"`
}

// Стоимость покупки value монет
func (c *SDK) EstimateCoinBuy(coinSmbl string, coinSmbl2 string, value int64) EstimateResponse {
	pip18 := Bip2Pip_i64(value).String() // монета в pip
	urlB := fmt.Sprintf("%s/api/estimateCoinBuy?coin_to_sell=%s&value_to_buy=%s&coin_to_buy=%s", c.MnAddress, coinSmbl2, pip18, coinSmbl)

	resB, err := http.Get(urlB)
	if err != nil {
		panic(err.Error())
	}
	defer resB.Body.Close()

	bodyB, err := ioutil.ReadAll(resB.Body)
	if err != nil {
		panic(err.Error())
	}

	var dataB node_estimate
	json.Unmarshal(bodyB, &dataB)

	return dataB.Result
}

// Стоимость продажи value монет
func (c *SDK) EstimateCoinSell(coinSmbl string, coinSmbl2 string, value int64) EstimateResponse {
	pip18 := Bip2Pip_i64(value).String() // монета в pip
	urlS := fmt.Sprintf("%s/api/estimateCoinSell?coin_to_sell=%s&value_to_sell=%s&coin_to_buy=%s", c.MnAddress, coinSmbl, pip18, coinSmbl2)

	resS, err := http.Get(urlS)
	if err != nil {
		panic(err.Error())
	}
	defer resS.Body.Close()

	bodyS, err := ioutil.ReadAll(resS.Body)
	if err != nil {
		panic(err.Error())
	}

	var dataS node_estimate
	json.Unmarshal(bodyS, &dataS)

	return dataS.Result
}
