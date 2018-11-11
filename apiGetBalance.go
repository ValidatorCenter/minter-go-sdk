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
// TODO: map[string]float32 <- PipStr2Bip_f32
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
