package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type node_status struct {
	Code   int
	Result result_node
}

// TODO: Преобразовать из string в int
type result_node struct {
	Version           string
	LatestBlockHash   string `json:"latest_block_hash"`
	LatestAppHash     string `json:"latest_app_hash"`
	LatestBlockHeight string `json:"latest_block_height"`
	LatestBlockTime   string `json:"latest_block_time"`
}

// получение сколько всего блоков в сети
func (c *SDK) GetStatus() result_node {
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
	/*latestBlcs, err := strconv.Atoi(data.Result.LatestBlockHeight)
	if err != nil {
		panic(err.Error())
	}*/

	return data.Result
}
