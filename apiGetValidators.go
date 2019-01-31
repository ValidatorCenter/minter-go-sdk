package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
	PubKey        string `json:"pubkey" bson:"pubkey" gorm:"pubkey" db:"pubkey"`
	VotingPowerTx string `json:"voting_power" bson:"-" gorm:"-" db:"-"`
	VotingPower   uint64 `json:"voting_power_u32" bson:"voting_power_u32" gorm:"voting_power_u32" db:"voting_power_u32"`
}

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
		VotingPower_i32, err := strconv.Atoi(data.Result[i1].VotingPowerTx)
		if err != nil {
			data.Result[i1].VotingPower = 0
		} else {
			data.Result[i1].VotingPower = uint64(VotingPower_i32)
		}
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
		VotingPower_i32, err := strconv.Atoi(data.Result[i1].VotingPowerTx)
		if err != nil {
			data.Result[i1].VotingPower = 0
		} else {
			data.Result[i1].VotingPower = uint64(VotingPower_i32)
		}
	}

	return data.Result, nil
}
