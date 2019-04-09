package mintersdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Результат выполнения получения баланас пользователя
type addrss_usr struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  AddrssResponse
	Error   ErrorStruct
}

type AddrssResponse struct {
	Balance            map[string]string `json:"balance" bson:"balance" gorm:"balance" db:"balance"`
	TransactionCountTx string            `json:"transaction_count" bson:"-" gorm:"-" db:"-"`
}

// узнаем баланс и количество транзакций
func (c *SDK) GetAddress(usrAddr string) (map[string]float32, uint32, error) {
	retDt := map[string]float32{}
	url := fmt.Sprintf("%s/address?address=%s", c.MnAddress, usrAddr)
	res, err := http.Get(url)
	if err != nil {
		return retDt, 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return retDt, 0, err
	}

	var data addrss_usr
	json.Unmarshal(body, &data)

	if data.Error.Code != 0 {
		err = errors.New(fmt.Sprint(data.Error.Code, " - ", data.Error.Message))
		return retDt, 0, err
	}

	for iS, vD := range data.Result.Balance {
		retDt[iS] = pipStr2bip_f32(vD)
	}

	txCount, err := strconv.Atoi(data.Result.TransactionCountTx)
	if err != nil {
		return retDt, 0, err
	}

	return retDt, uint32(txCount), nil
}
