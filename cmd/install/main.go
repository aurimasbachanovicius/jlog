package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/aurimasbachanovicius/jlog/v2/pkg/config"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("could not get current os user: %s", err)
	}

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

	err = config.CreateConfig(usr.HomeDir+"/.jlog/config/", "config.yaml", config.Config{
		JiraServerDomain: domain,
		JiraAuth:         fmt.Sprintf("Basic %s", auth),
	})

	if err != nil {
		log.Fatalf("Could not create config: %s", err)
	}
}
