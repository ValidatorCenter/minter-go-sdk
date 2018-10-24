package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Результат выполнения получения номера операции
type count_transaction struct {
	Code   int                `json:"code"`
	Result TransCountResponse `json:"result"`
}
type TransCountResponse struct {
	Count int `json:"count"`
}

// Возвращает количество исходящих транзакций с данной учетной записи.
func (c *SDK) GetNonce(txAddress string) int {
	url := fmt.Sprintf("%s/api/transactionCount/%s", c.MnAddress, txAddress)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data count_transaction
	json.Unmarshal(body, &data)
	return data.Result.Count
}
