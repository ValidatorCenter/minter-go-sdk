package mintersdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// запрос по кандидату (curl -s 'localhost:8841/api/candidate/{public_key}')
type node_candidate struct {
	Code   int
	Result result_candidate
}

type result_candidate struct {
	Candidate candidate_info `json:"candidate" bson:"candidate" gorm:"candidate"`
}

// type candidate_info struct --- в apiGetCandidates.go

func (c *SDK) GetCandidate(candidateHash string) candidate_info {
	url := fmt.Sprintf("%s/api/candidate/%s", c.MnAddress, candidateHash)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data node_candidate
	json.Unmarshal(body, &data)
	return data.Result.Candidate
}
