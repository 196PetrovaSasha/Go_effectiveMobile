package create

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type NationalizeData struct {
	Name         string              `json:"name"`
	CountryCodes []CountryCodeResult `json:"country"`
}

type CountryCodeResult struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

func GetNationality(name string) (string, error) {
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return "", nil
	}

	var data NationalizeData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return "", err
	}

	return data.CountryCodes[0].CountryID, nil
}
