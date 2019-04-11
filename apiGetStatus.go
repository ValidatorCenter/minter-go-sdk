package mintersdk

import (
	//"encoding/json" -- переход на easyjson
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//easyjson:json
type node_status struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Result  ResultNetwork `json:"result"`
	Error   ErrorStruct   `json:"error"`
}

type ResultNetwork struct {
	Version             string    `json:"version" bson:"version" gorm:"version" db:"version"`
	LatestBlockHash     string    `json:"latest_block_hash" bson:"-" gorm:"-" db:"-"`
	LatestAppHash       string    `json:"latest_app_hash" bson:"-" gorm:"-" db:"-"`
	LatestBlockHeightTx string    `json:"latest_block_height" bson:"-" gorm:"-" db:"-"`
	LatestBlockHeight   int       `json:"latest_block_height_i32" bson:"latest_block_height_i32" gorm:"latest_block_height_i32" db:"latest_block_height_i32"`
	LatestBlockTime     time.Time `json:"latest_block_time" bson:"-" gorm:"-" db:"-"`
	// tm_status {...}
}

// получение сколько всего блоков в сети
func (c *SDK) GetStatus() (ResultNetwork, error) {
	url := fmt.Sprintf("%s/status", c.MnAddress)
	res, err := http.Get(url)
	if err != nil {
		return ResultNetwork{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ResultNetwork{}, err
	}

	var data node_status
	//json.Unmarshal(body, &data) -- переход на easyjson

	err = data.UnmarshalJSON(body)
	if err != nil {
		panic(err)
	}

	if data.Error.Code != 0 {
		err = errors.New(fmt.Sprint(data.Error.Code, " - ", data.Error.Message))
		return ResultNetwork{}, err
	}

	data.Result.LatestBlockHeight, err = strconv.Atoi(data.Result.LatestBlockHeightTx)
	if err != nil {
		// пусть и не полностью
		return data.Result, err
	}

	return data.Result, nil
}
