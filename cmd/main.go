package main

import (
	"context"
	"os"
	"regexp"
	"strings"

	"github.com/google/go-github/v32/github"
)

const maxRepositoriesCount = 256 // this is max value for github job schedulers

/*
	os.Args[1] - regex pattern
	os.Args[2] - org name
*/
func main() {
	client := github.NewClient(nil)
	opt := &github.RepositoryListByOrgOptions{Type: "public", ListOptions: github.ListOptions{PerPage: maxRepositoriesCount}}
	if len(os.Args) != 3 {
		panic("args len is not 3")
	}
	regex := regexp.MustCompile(os.Args[1])
	orgName := os.Args[2]
	repos, _, err := client.Repositories.ListByOrg(context.Background(), orgName, opt)
	if err != nil {
		panic(err.Error())
	}
	var result []*github.Repository
	for _, r := range repos {
		if !regex.MatchString(r.GetName()) {
			continue
		}
		result = append(result, r)
	}

	var sb strings.Builder

	_, _ = sb.WriteString("[")
	for i, r := range result {
		_, _ = sb.WriteString("\"")
		_, _ = sb.WriteString(r.GetFullName())
		_, _ = sb.WriteString("\"")
		if i+1 < len(result) {
			_, _ = sb.WriteString(", ")
		}
	}
	_, _ = sb.WriteString("]")

	println("::set-output name=repositories::" + sb.String())
}
