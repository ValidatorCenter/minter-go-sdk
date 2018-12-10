package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// запрос по валидаторам
type node_validators struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  []result_valid
	Error   ErrorStruct
}

// результат по валидаторам
type result_valid struct {
	AccumulatedRewardTx string        `json:"accumulated_reward" bson:"-" gorm:"-"`
	AccumulatedReward   float32       `json:"accumulated_reward_f32" bson:"accumulated_reward_f32" gorm:"accumulated_reward_f32"`
	AbsentTimes         int           `json:"absent_times" bson:"absent_times" gorm:"absent_times"`
	Candidate           CandidateInfo `json:"candidate" bson:"candidate" gorm:"candidate"`
}

// type CandidateInfo struct --- в apiGetCandidates.go

// Возвращает список валидаторов по номеру блока (у мастерноды должен быть включен keep_state_history)
func (c *SDK) GetValidatorsBlock(blockN int) ([]result_valid, error) {
	url := fmt.Sprintf("%s/validators?height=%d", c.MnAddress, blockN)
	res, err := http.Get(url)
	if err != nil {
		return []result_valid{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []result_valid{}, err
	}

	var data node_validators
	json.Unmarshal(body, &data)

	for i1, _ := range data.Result {
		data.Result[i1].AccumulatedReward = pipStr2bip_f32(data.Result[i1].AccumulatedRewardTx)
		data.Result[i1].Candidate.TotalStake = pipStr2bip_f32(data.Result[i1].Candidate.TotalStakeTx)
	}

	return data.Result, nil
}

// Возвращает список валидаторов
func (c *SDK) GetValidators() ([]result_valid, error) {
	url := fmt.Sprintf("%s/validators", c.MnAddress)
	res, err := http.Get(url)
	if err != nil {
		return []result_valid{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []result_valid{}, err
	}

	var data node_validators
	json.Unmarshal(body, &data)

	for i1, _ := range data.Result {
		data.Result[i1].AccumulatedReward = pipStr2bip_f32(data.Result[i1].AccumulatedRewardTx)
		data.Result[i1].Candidate.TotalStake = pipStr2bip_f32(data.Result[i1].Candidate.TotalStakeTx)
		// В новом API нет у "candidates" Стэка!!!
	}

	return data.Result, nil
}
