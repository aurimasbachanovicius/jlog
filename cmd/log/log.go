package main

import (
	"fmt"
)

func (a app) log(time string, key string) error {
	err := a.jira.Log(time, key)
	if err != nil {
		return fmt.Errorf("could not jira log: %s", err)
	}

	return nil
}
