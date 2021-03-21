package main

import "os"

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
	if auth == "" {
		auth, domain = install()
		err := os.Setenv("JLOG_AUTH", auth)
		if err != nil {
			panic("Could not set JLOG_AUTH env. variable")
		}
		err = os.Setenv("JLOG_DOMAIN", domain)
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
