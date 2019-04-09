package mintersdk

import (
	"encoding/json"
	"errors"
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
	Events []BlockEventsResponse `json:"events" bson:"events" gorm:"events" db:"events"`
}

type BlockEventsResponse struct {
	Type  string         `json:"type" bson:"type" gorm:"type" db:"type"`
	Value EventValueData `json:"value" bson:"value" gorm:"value" db:"value"`
}

type EventValueData struct {
	Role            string  `json:"role" bson:"role" gorm:"role" db:"role"` //DAO,Developers,Validator,Delegator
	Address         string  `json:"address" bson:"address" gorm:"address" db:"address"`
	AmountTx        string  `json:"amount" bson:"-" gorm:"-" db:"-"`
	Amount          float32 `json:"amount_f32" bson:"amount_f32" gorm:"amount_f32" db:"amount_f32"`
	Coin            string  `json:"coin" bson:"coin" gorm:"coin" db:"coin"`
	ValidatorPubKey string  `json:"validator_pub_key" bson:"validator_pub_key" gorm:"validator_pub_key" db:"validator_pub_key"`
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

	if data.Error.Code != 0 {
		err = errors.New(fmt.Sprint(data.Error.Code, " - ", data.Error.Message))
		return BlockEvResponse{}, err
	}

	for iStep, _ := range data.Result.Events {
		data.Result.Events[iStep].Value.Amount = pipStr2bip_f32(data.Result.Events[iStep].Value.AmountTx)
		if data.Result.Events[iStep].Value.Coin == "" {
			data.Result.Events[iStep].Value.Coin = GetBaseCoin()
		}
		/*fmt.Printf("DEFCOIN: %s\n", GetBaseCoin())
		fmt.Printf("%#v\n", data.Result.Events[iStep].Value)*/
	}

	return data.Result, nil
}
