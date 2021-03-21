package main

import (
	"github.com/aurimasbachanovicius/jlog/v2/pkg/jira"
)

type app struct {
	jira *jira.Jira
}

func main() {
	args := resolveArgs()

	app := app{
		jira: &jira.Jira{
			Auth:   args.auth,
			Domain: args.domain,
		},
	}

	err := app.log(args.time, args.key)
	if err != nil {
		panic(err)
	}
}
