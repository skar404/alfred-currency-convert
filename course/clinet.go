package course

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Item struct {
	Name     string  `json:"Name"`
	CharCode string  `json:"CharCode"`
	Value    float32 `json:"Value"`
}

type CBRRequest struct {
	Date time.Time       `json:"Date"`
	Data map[string]Item `json:"Valute"`
}

type CBRData struct {
	Date time.Time
	Data []Item
}

func jsonHttpClient(method string, url string, body, target interface{}) {
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		panic(err)
	}
}

//GetCBRCourse
//get data in user api: https://www.cbr-xml-daily.ru/
func getRawCBRCourse() CBRRequest {
	var reqData CBRRequest

	jsonHttpClient("GET", "https://www.cbr-xml-daily.ru/daily_json.js", &CBRRequest{}, &reqData)
	return reqData
}

func GetCBRCourse() CBRData {
	var dataList []Item

	req := getRawCBRCourse()
	for _, element := range req.Data {
		dataList = append(dataList, element)
	}

	return CBRData{
		Date: req.Date,
		Data: dataList,
	}
}
