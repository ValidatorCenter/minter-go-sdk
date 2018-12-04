package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type node_estimate struct {
	Code   int
	Result _estimateResponse
}

type _estimateResponse struct {
	WillPay    string `json:"will_pay"`
	WillGet    string `json:"will_get"`
	Commission string `json:"commission"`
}

// для запроса о стоимости монет
type EstimateResponse struct {
	WillPay    float32 `json:"will_pay" bson:"will_pay" gorm:"will_pay"`
	WillGet    float32 `json:"will_get" bson:"will_get" gorm:"will_get"`
	Commission float32 `json:"commission" bson:"commission" gorm:"commission"`
}

// Стоимость покупки value монет
func (c *SDK) EstimateCoinBuy(coinBuy string, coinSell string, valueBuy int64) (EstimateResponse, error) {
	pip18 := bip2pip_i64(valueBuy).String() // монета в pip
	urlB := fmt.Sprintf("%s/api/estimateCoinBuy?coin_to_sell=%s&value_to_buy=%s&coin_to_buy=%s", c.MnAddress, coinSell, pip18, coinBuy)

	resB, err := http.Get(urlB)
	if err != nil {
		return EstimateResponse{}, err
	}
	defer resB.Body.Close()

	bodyB, err := ioutil.ReadAll(resB.Body)
	if err != nil {
		return EstimateResponse{}, err
	}

	var dataB node_estimate
	json.Unmarshal(bodyB, &dataB)

	return EstimateResponse{
		WillPay:    pipStr2bip_f32(dataB.Result.WillPay),
		WillGet:    pipStr2bip_f32(dataB.Result.WillGet),
		Commission: pipStr2bip_f32(dataB.Result.Commission),
	}, nil
}

// Стоимость продажи value монет
func (c *SDK) EstimateCoinSell(coinSell string, coinBuy string, valueSell int64) (EstimateResponse, error) {
	pip18 := bip2pip_i64(valueSell).String() // монета в pip
	urlS := fmt.Sprintf("%s/api/estimateCoinSell?coin_to_sell=%s&value_to_sell=%s&coin_to_buy=%s", c.MnAddress, coinSell, pip18, coinBuy)

	resS, err := http.Get(urlS)
	if err != nil {
		return EstimateResponse{}, err
	}
	defer resS.Body.Close()

	bodyS, err := ioutil.ReadAll(resS.Body)
	if err != nil {
		return EstimateResponse{}, err
	}

	var dataS node_estimate
	json.Unmarshal(bodyS, &dataS)

	return EstimateResponse{
		WillPay:    pipStr2bip_f32(dataS.Result.WillPay),
		WillGet:    pipStr2bip_f32(dataS.Result.WillGet),
		Commission: pipStr2bip_f32(dataS.Result.Commission),
	}, nil
}
