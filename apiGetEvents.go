package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Содержимое блока
type node_block_ev struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  BlockEvResponse
	Error   ErrorStruct
}

type BlockEvResponse struct {
	Events []BlockEventsResponse `json:"events" bson:"events" gorm:"events"`
}

type BlockEventsResponse struct {
	Type  string         `json:"type" bson:"type" gorm:"type"`
	Value EventValueData `json:"value" bson:"value" gorm:"value"`
}

type EventValueData struct {
	Role            string  `json:"role" bson:"role" gorm:"role"` //DAO,Developers,Validator,Delegator
	Address         string  `json:"address" bson:"address" gorm:"address"`
	AmountTx        string  `json:"amount" bson:"-" gorm:"-"`
	Amount          float32 `json:"amount_f32" bson:"amount_f32" gorm:"amount_f32"`
	ValidatorPubKey string  `json:"validator_pub_key" bson:"validator_pub_key" gorm:"validator_pub_key"`
}

// получаем содержимое событий блока по его ID
func (c *SDK) GetEvents(id int) (BlockEvResponse, error) {
	url := fmt.Sprintf("%s/events?height=%d", c.MnAddress, id)
	res, err := http.Get(url)
	if err != nil {
		return BlockEvResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return BlockEvResponse{}, err
	}

	var data node_block_ev
	json.Unmarshal(body, &data)

	for iStep, _ := range data.Result.Events {
		data.Result.Events[iStep].Value.Amount = pipStr2bip_f32(data.Result.Events[iStep].Value.AmountTx)
	}

	return data.Result, nil
}
