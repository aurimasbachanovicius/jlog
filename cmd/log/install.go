package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/aurimasbachanovicius/jlog/v2/pkg/config"
)

func install(dir string, file string) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Jira server domain -> ")
	domain, _ := reader.ReadString('\n')
	domain = strings.Replace(domain, "\n", "", -1)

	fmt.Print("Jira Username (email) -> ")
	email, _ := reader.ReadString('\n')
	email = strings.Replace(email, "\n", "", -1)

	fmt.Print("Jira API token -> ")
	token, _ := reader.ReadString('\n')
	token = strings.Replace(token, "\n", "", -1)

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", email, token)))

	err := config.CreateConfig(dir, file, config.Config{
		JiraServerDomain: domain,
		JiraAuth:         fmt.Sprintf("Basic %s", auth),
	})

	if err != nil {
		return fmt.Errorf("could not create config: %s", err)
	}

	return nil
}
