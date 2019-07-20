package course

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type SaveItem struct {
	Code  string  `json:"code"`
	Value float32 `json:"value"`
}

type SaveRequest struct {
	Date time.Time  `json:"time"`
	Data []SaveItem `json:"data"`
}

type SaveData struct {
	Follows  []string    `json:"follows"`
	SaveData SaveRequest `json:"save_data"`
}

func (r *SaveData) Load() error {
	plan, _ := ioutil.ReadFile("save.json")

	err := json.Unmarshal(plan, &r)
	return err
}

func (r *SaveRequest) Save() error {
	saveJson, _ := json.Marshal(r)
	err := ioutil.WriteFile("save.json", saveJson, 0644)
	return err
}
