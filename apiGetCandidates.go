package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// запрос на всех кандидатов (curl -s 'localhost:8841/api/candidates')
type node_candidates struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  []CandidateInfo
	Error   ErrorStruct
}

// type CandidateInfo struct --- в apiGetCandidate.go

// Возвращает список нод валидаторов и кандидатов
func (c *SDK) GetCandidates() ([]CandidateInfo, error) {
	url := fmt.Sprintf("%s/candidates", c.MnAddress)
	res, err := http.Get(url)
	if err != nil {
		return []CandidateInfo{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []CandidateInfo{}, err
	}

	var data node_candidates
	json.Unmarshal(body, &data)

	for i1, _ := range data.Result {
		data.Result[i1].TotalStake = pipStr2bip_f32(data.Result[i1].TotalStakeTx)
		// В новом API нет у "candidates" Стэка!!!
		/*for i2, _ := range data.Result[i1].Stakes {
			data.Result[i1].Stakes[i2].Value = pipStr2bip_f32(data.Result[i1].Stakes[i2].ValueTx)
			data.Result[i1].Stakes[i2].BipValue = pipStr2bip_f32(data.Result[i1].Stakes[i2].BipValueTx)
		}*/
	}
	return data.Result, nil
}
