package smsservice

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	URL           = "https://www.fast2sms.com/dev/bulkV2"
	Authorization = "zEvho6AFONMnmPB1RLgJDayw8iSVbeG4qlsUZQcTCpuXKY95dk3yxoYzJLPb0IOrvKtpHMnmWcEG1f76"
)

func SendOTP(toNumbers []string, otps []string) (map[string]interface{}, error) {
	client := &http.Client{}

	payload := strings.NewReader(`{
		"route": "dlt",
		"sender_id": "TWOBTS",
		"message": "130138",
		"variables_values":"` + strings.Join(otps, ",") + `",
		"flash": 0,
		"numbers":"` + strings.Join(toNumbers, ",") + `"
	}`)

	req, err := http.NewRequest("POST", URL, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("authorization", Authorization)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
