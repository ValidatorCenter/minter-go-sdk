package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type node_status struct {
	Code   int
	Result ResultNetwork
}

type ResultNetwork struct {
	Version             string
	LatestBlockHash     string    `json:"latest_block_hash" bson:"-" gorm:"-"`
	LatestAppHash       string    `json:"latest_app_hash" bson:"-" gorm:"-"`
	LatestBlockHeightTx string    `json:"latest_block_height" bson:"-" gorm:"-"`
	LatestBlockHeight   int       `json:"latest_block_height_i32" bson:"latest_block_height_i32" gorm:"latest_block_height_i32"`
	LatestBlockTime     time.Time `json:"latest_block_time" bson:"-" gorm:"-"`
	// tm_status {...}
}

// получение сколько всего блоков в сети
func (c *SDK) GetStatus() ResultNetwork {
	url := fmt.Sprintf("%s/api/status", c.MnAddress)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data node_status
	json.Unmarshal(body, &data)

	data.Result.LatestBlockHeight, err = strconv.Atoi(data.Result.LatestBlockHeightTx)
	if err != nil {
		panic(err.Error())
	}

	return data.Result
}
