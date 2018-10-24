package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// запрос на всех кандидатов (curl -s 'localhost:8841/api/candidates')
type node_candidates struct {
	Code   int
	Result []candidate_info
}

// структура кандидата/валидатора
type candidate_info struct {
	CandidateAddress string `json:"candidate_address"`
	TotalStake       string `json:"total_stake"`
	TotalStake32     float32
	PubKey           string        `json:"pub_key"`
	Commission       int           `json:"commission"`
	CreatedAtBlock   int           `json:"created_at_block"`
	StatusInt        int           `json:"status"` // числовое значение статуса: 1 - Offline, 2 - Online
	Stakes           []stakes_info `json:"stakes"`
}

// стэк делегатов
type stakes_info struct {
	Owner      string `json:"owner"`
	Coin       string `json:"coin"`
	Value      string `json:"value"`
	BipValue   string `json:"bip_value"`
	Value32    float32
	BipValue32 float32
}

// Возвращает список нод валидаторов и кандидатов
func (c *SDK) GetCandidates() []candidate_info {
	url := fmt.Sprintf("%s/api/candidates", c.MnAddress)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data node_candidates
	json.Unmarshal(body, &data)
	return data.Result
}
