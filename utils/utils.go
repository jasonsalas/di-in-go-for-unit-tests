package utils

import (
	"encoding/json"

	"github.com/jasonsalas/di-in-go-for-unit-tests/astros"
)

const (
	API = "http://api.open-notify.org/astros.json"
)

func GetAstronauts(getWebRequest astros.GetWebRequest) (int, error) {
	body, err := getWebRequest.FetchBytes(API)
	if err != nil {
		return 0, err
	}

	peopleResult := astros.People{}

	if err := json.Unmarshal(body, &peopleResult); err != nil {
		return 0, err
	}
	return peopleResult.Number, err
}
