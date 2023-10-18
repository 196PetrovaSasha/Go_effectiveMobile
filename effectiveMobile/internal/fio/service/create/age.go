package create

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AgifyData struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}

func GetAge(name string) (int, error) {
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)

	response, err := http.Get(url)
	if err != nil {
		return -1, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return -1, err
	}

	var data AgifyData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return -1, err
	}

	return data.Age, nil
}
