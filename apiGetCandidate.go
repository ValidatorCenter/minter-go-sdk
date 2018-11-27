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
	Candidate CandidateInfo `json:"candidate" bson:"candidate" gorm:"candidate"`
}

// type CandidateInfo struct --- в apiGetCandidates.go

func (c *SDK) GetCandidate(candidateHash string) (CandidateInfo, error) {
	url := fmt.Sprintf("%s/api/candidate/%s", c.MnAddress, candidateHash)
	res, err := http.Get(url)
	if err != nil {
		return CandidateInfo{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return CandidateInfo{}, err
	}

	var data node_candidate
	json.Unmarshal(body, &data)

	data.Result.Candidate.TotalStake = pipStr2bip_f32(data.Result.Candidate.TotalStakeTx)
	for i2, _ := range data.Result.Candidate.Stakes {
		data.Result.Candidate.Stakes[i2].Value = pipStr2bip_f32(data.Result.Candidate.Stakes[i2].ValueTx)
		data.Result.Candidate.Stakes[i2].BipValue = pipStr2bip_f32(data.Result.Candidate.Stakes[i2].BipValueTx)
	}

	return data.Result.Candidate, nil
}
