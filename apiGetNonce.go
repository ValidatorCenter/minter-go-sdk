package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Результат выполнения получения номера операции
type count_transaction struct {
	Code   int
	Result TransCountResponse
}
type TransCountResponse struct {
	Count int `json:"count" bson:"count" gorm:"count"`
}

// Возвращает количество исходящих транзакций с данной учетной записи.
func (c *SDK) GetNonce(txAddress string) (int, error) {
	url := fmt.Sprintf("%s/api/transactionCount/%s", c.MnAddress, txAddress)
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	var data count_transaction
	json.Unmarshal(body, &data)
	return data.Result.Count, nil
}
