package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	t "time"
)

type worklog struct {
	TimeSpentSeconds int `json:"timeSpentSeconds"`
	Visibility       struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"visibility"`
	Comment struct {
		Type    string `json:"type"`
		Version int    `json:"version"`
		Content []struct {
			Type    string `json:"type"`
			Content []struct {
				Text string `json:"text"`
				Type string `json:"type"`
			} `json:"content"`
		} `json:"content"`
	} `json:"comment"`
	Started string `json:"started"`
}

func (j Jira) Log(time int, key string) error {
	url := fmt.Sprintf("https://%s/rest/api/3/issue/%s/worklog", j.Domain, key)

	jsonData, err := json.Marshal(worklog{
		TimeSpentSeconds: time,
		Started:          t.Now().UTC().Format("2006-01-02T15:04:05-0700"),
	})

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("could not make error request: %s", err)
	}

	req.Header.Add("Authorization", j.Auth)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not client do: %s", err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read: %s", err)
	}

	return nil
}
