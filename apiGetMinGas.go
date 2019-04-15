package mintersdk

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//easyjson:json
type min_gas struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Result  int64       `json:"result,string"`
	Error   ErrorStruct `json:"error"`
}

// получение минимального значения газа на данный момент
func (c *SDK) GetMinGas() (int64, error) {
	url := fmt.Sprintf("%s/min_gas_price", c.MnAddress)
	res, err := http.Get(url)
	if err != nil {
		return 1, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 1, err
	}

	var data min_gas

	err = data.UnmarshalJSON(body)
	if err != nil {
		panic(err)
	}

	if data.Error.Code != 0 {
		err = errors.New(fmt.Sprint(data.Error.Code, " - ", data.Error.Message))
		return 1, err
	}

	return data.Result, nil
}
