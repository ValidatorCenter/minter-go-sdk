package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bytes"
	"encoding/hex"
	"errors"

	tr "github.com/MinterTeam/minter-go-node/core/transaction"
)

// Ответ транзакции
type send_transaction struct {
	Code   int
	Result TransSendResponse
	Log    string
}
type TransSendResponse struct {
	Hash string `json:"hash" bson:"hash" gorm:"hash"`
}

// Исполнение транзакции закодированной RLP
func (c *SDK) SetTransaction(tx *tr.Transaction) (string, error) {

	encodedTx, err := tx.Serialize()
	if err != nil {
		panic(err)
	}

	strTxRPL := hex.EncodeToString(encodedTx)

	message := map[string]interface{}{
		"transaction": strTxRPL,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		fmt.Println("ERROR: SetCandidateTransaction::json.Marshal")
		return "", err
	}

	url := fmt.Sprintf("%s/api/sendTransaction", c.MnAddress)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		fmt.Println("ERROR: TxSign::http.Post")
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERROR: TxSign::ioutil.ReadAll")
		return "", err
	}

	var data send_transaction
	json.Unmarshal(body, &data)

	if data.Code == 0 {
		return data.Result.Hash, nil
	} else {
		fmt.Printf("ERROR: TxSign: %#v\n", data)
		return data.Log, errors.New(fmt.Sprintf("Err:%d", data.Code))
	}
}
