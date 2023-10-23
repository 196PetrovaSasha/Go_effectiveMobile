package create

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GenderizeData struct {
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
	Count       int     `json:"count"`
}

func GetGender(name string) (string, error) {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var data GenderizeData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	return data.Gender, nil
}
