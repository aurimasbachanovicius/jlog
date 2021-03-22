package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/aurimasbachanovicius/jlog/v2/pkg/config"
	"github.com/aurimasbachanovicius/jlog/v2/pkg/jira"
)

type app struct {
	jira *jira.Jira
}

type args struct {
	time string
	key  string
}

func resolveArgs() *args {
	if len(os.Args) <= 2 {
		log.Fatal("usage: log <task> <time>")
	}

	return &args{
		key:  os.Args[1],
		time: os.Args[2],
	}
}

func resolveConfig() *config.Config {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("could not get current os user: %s", err)
	}

	c, err := config.GetConfig(usr.HomeDir + "/.jlog/config/config.yaml")
	if err != nil {
		log.Fatalf("could not get configs: %s", err)
	}

	return c
}

func main() {
	args := resolveArgs()
	conf := resolveConfig()

	app := app{
		jira: &jira.Jira{
			Auth:   conf.JiraAuth,
			Domain: conf.JiraServerDomain,
		},
	}

	err := app.log(args.time, args.key)
	if err != nil {
		log.Fatal(fmt.Errorf("app could not log: %s", err))
	}
}
