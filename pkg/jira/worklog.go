package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	t "time"
)

type contentMeta struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type content struct {
	Type    string        `json:"type"`
	Content []contentMeta `json:"content"`
}

type comment struct {
	Type    string    `json:"type"`
	Version int       `json:"version"`
	Content []content `json:"content"`
}

type worklog struct {
	TimeSpentSeconds int      `json:"timeSpentSeconds"`
	Visibility       struct{} `json:"visibility"`
	Comment          comment  `json:"comment"`
	Started          string   `json:"started"`
}

func (j Jira) Log(time int, key string) error {
	url := fmt.Sprintf("https://%s/rest/api/3/issue/%s/worklog", j.Domain, key)

	jsonData, err := json.Marshal(worklog{
		TimeSpentSeconds: time,
		Started:          t.Now().UTC().Format("2006-01-02T15:04:05.999-0700"),
		Comment: comment{
			Type:    "doc",
			Version: 1,
			Content: []content{
				{
					Type: "paragraph",
					Content: []contentMeta{{
						Text: "jlog logged",
						Type: "text",
					}},
				},
			},
		},
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
