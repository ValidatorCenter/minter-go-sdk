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
	Result []CandidateInfo
}

// структура кандидата/валидатора (экспортная)
type CandidateInfo struct {
	CandidateAddress string        `json:"candidate_address" bson:"candidate_address" gorm:"candidate_address"`
	TotalStakeTx     string        `json:"total_stake" bson:"-" gorm:"-"`
	TotalStake       float32       `json:"total_stake_f32" bson:"total_stake_f32" gorm:"total_stake_f32"`
	PubKey           string        `json:"pub_key" bson:"pub_key" gorm:"pub_key"`
	Commission       int           `json:"commission" bson:"commission" gorm:"commission"`
	CreatedAtBlock   int           `json:"created_at_block" bson:"created_at_block" gorm:"created_at_block"`
	StatusInt        int           `json:"status" bson:"status" gorm:"status"` // числовое значение статуса: 1 - Offline, 2 - Online
	Stakes           []stakes_info `json:"stakes" bson:"stakes" gorm:"stakes"`
}

// стэк делегатов
type stakes_info struct {
	Owner      string  `json:"owner" bson:"owner" gorm:"owner"`
	Coin       string  `json:"coin" bson:"coin" gorm:"coin"`
	ValueTx    string  `json:"value" bson:"-" gorm:"-"`
	BipValueTx string  `json:"bip_value" bson:"-" gorm:"-"`
	Value      float32 `json:"value_f32" bson:"value_f32" gorm:"value_f32"`
	BipValue   float32 `json:"bip_value_f32" bson:"bip_value_f32" gorm:"bip_value_f32"`
}

// Возвращает список нод валидаторов и кандидатов
func (c *SDK) GetCandidates() []CandidateInfo {
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

	for i1, _ := range data.Result {
		data.Result[i1].TotalStake = pipStr2bip_f32(data.Result[i1].TotalStakeTx)
		for i2, _ := range data.Result[i1].Stakes {
			data.Result[i1].Stakes[i2].Value = pipStr2bip_f32(data.Result[i1].Stakes[i2].ValueTx)
			data.Result[i1].Stakes[i2].BipValue = pipStr2bip_f32(data.Result[i1].Stakes[i2].BipValueTx)
		}
	}
	return data.Result
}
