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

func resolveConfig(dir string, file string) *config.Config {
	c, err := config.GetConfig(dir + file)
	if err != nil {
		log.Fatalf("could not get configs: %s", err)
	}

	return c
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("could not get current os user: %s", err)
	}

	dir := usr.HomeDir + "/.jlog/config/"
	file := "config.yaml"

	if os.Args[1] == "install" {
		err = install(dir, file)

		if err != nil {
			log.Fatalf("Could not create config: %s", err)
		}
	}

	args := resolveArgs()
	conf := resolveConfig(dir, file)

	app := app{
		jira: &jira.Jira{
			Auth:   conf.JiraAuth,
			Domain: conf.JiraServerDomain,
		},
	}

	err = app.log(args.time, args.key)
	if err != nil {
		log.Fatal(fmt.Errorf("app could not log: %s", err))
	}
}
