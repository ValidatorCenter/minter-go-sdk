package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Результат выполнения получения баланас пользователя
type blnc_usr struct {
	Code   int           `json:"code"`
	Result BlnctResponse `json:"result"`
}
type BlnctResponse struct {
	Balance map[string]string `json:"balance"`
}

// узнаем баланс
// TODO: map[string]float32 <- cnvStr2Float_18
func (c *SDK) GetBalance(usrAddr string) map[string]string {
	url := fmt.Sprintf("%s/api/balance/%s", c.MnAddress, usrAddr)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data blnc_usr
	json.Unmarshal(body, &data)

	return data.Result.Balance
}
