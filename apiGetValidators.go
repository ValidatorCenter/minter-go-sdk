package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// запрос по валидаторам
type node_validators struct {
	Code   int
	Result []result_valid
}

// результат по валидаторам
type result_valid struct {
	AccumulatedReward   string `json:"accumulated_reward"`
	AccumulatedReward32 float32
	AbsentTimes         int `json:"absent_times"`
	Candidate           candidate_info
}

// type candidate_info struct --- в apiGetCandidates.go

// TODO: Возвращает список валидаторов по номеру блока

// Возвращает список валидаторов
func (c *SDK) GetValidators() []result_valid {
	url := fmt.Sprintf("%s/api/validators", c.MnAddress)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data node_validators
	json.Unmarshal(body, &data)
	return data.Result
}
