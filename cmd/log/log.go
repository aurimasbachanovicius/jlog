package main

import (
	"fmt"
	"strconv"

	timeconv "github.com/aurimasbachanovicius/jlog/v2/pkg/time"
)

func (a app) log(time string, key string) error {
	minutes, err := strconv.Atoi(time)
	if err != nil {
		return fmt.Errorf("could not convert string to int: %s", err)
	}

	err = a.jira.Log(timeconv.FromMinToSec(minutes), key)
	if err != nil {
		return fmt.Errorf("could not jira log: %s", err)
	}

	return nil
}
