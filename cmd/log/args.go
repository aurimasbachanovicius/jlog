package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type args struct {
	time   string
	key    string
	auth   string
	domain string
}

func resolveArgs() args {
	if len(os.Args) <= 2 {
		panic("usage: log <task> <time>")
	}

	auth := os.Getenv("JLOG_AUTH")
	domain := os.Getenv("JLOG_DOMAIN")

	reader := bufio.NewReader(os.Stdin)

	if auth == "" {
		fmt.Print("Jira API Token -> ")
		auth, _ := reader.ReadString('\n')
		auth = strings.Replace(auth, "\n", "", -1)
		err := os.Setenv("JLOG_AUTH", fmt.Sprintf("Basic %s", auth))
		if err != nil {
			panic("Could not set JLOG_AUTH env. variable")
		}
	}

	if domain == "" {
		fmt.Print("Jira Server domain -> ")
		domain, _ := reader.ReadString('\n')
		domain = strings.Replace(domain, "\n", "", -1)
		err := os.Setenv("JLOG_DOMAIN", domain)
		if err != nil {
			panic("Could not set JLOG_DOMAIN env. variable")
		}
	}

	return args{
		key:    os.Args[1],
		time:   os.Args[2],
		auth:   auth,
		domain: domain,
	}
}
