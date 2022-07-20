package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type JiraIssue struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
	} `json:"fields"`
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: jira-branch <url>")
	}
	username := os.Getenv("JIRA_USERNAME")
	password := os.Getenv("JIRA_API_KEY")

	if username == "" || password == "" {
		log.Fatal("Set JIRA_USERNAME and JIRA_API_KEY")
	}

	url := os.Args[1]

	match := regexp.MustCompile(`https?://(?P<subdomain>.*)/browse/(?P<ticket>\w+-\d+)$`).FindStringSubmatch(url)
	endpoint := fmt.Sprintf("https://%s/rest/api/3/issue/%s?fields=summary", match[1], match[2])

	client := &http.Client{}
	request, err := http.NewRequest("GET", endpoint, nil)
	check(err)
	request.SetBasicAuth(username, password)
	response, err := client.Do(request)
	check(err)
	if response.StatusCode != 200 {
		log.Fatalf("Unable to fetch issue data: %s", response.Status)
	}
	bodyText, err := ioutil.ReadAll(response.Body)
	check(err)

	var issue JiraIssue
	err = json.Unmarshal(bodyText, &issue)
	check(err)

	branchName := fmt.Sprintf(
		"%s/%s",
		issue.Key,
		regexp.MustCompile(`(?m)\W+`).ReplaceAllString(strings.ToLower(issue.Fields.Summary), "_"),
	)

	stdout, err := exec.Command("git", "checkout", branchName).CombinedOutput()
	if err != nil {
		stdout, err = exec.Command("git", "checkout", "-b", branchName).CombinedOutput()
		check(err)
		fmt.Print(string(stdout))

	} else {
		fmt.Print(string(stdout))
	}
}
