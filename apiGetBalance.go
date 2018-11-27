package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Результат выполнения получения баланас пользователя
type blnc_usr struct {
	Code   int
	Result BlnctResponse
}
type BlnctResponse struct {
	Balance map[string]string `json:"balance" bson:"balance" gorm:"balance"`
}

// узнаем баланс
func (c *SDK) GetBalance(usrAddr string) (map[string]float32, error) {
	url := fmt.Sprintf("%s/api/balance/%s", c.MnAddress, usrAddr)
	retDt := map[string]float32{}
	res, err := http.Get(url)
	if err != nil {
		return retDt, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return retDt, err
	}

	var data blnc_usr
	json.Unmarshal(body, &data)

	for iS, vD := range data.Result.Balance {
		retDt[iS] = pipStr2bip_f32(vD)
	}

	return retDt, nil
}
