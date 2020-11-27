package main

import (
	"context"
	"os"
	"strings"

	"github.com/google/go-github/v32/github"
)

func main() {
	client := github.NewClient(nil)
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), os.Args[1], opt)
	if err != nil {
		panic(err.Error())
	}
	var sb strings.Builder
	_, _ = sb.WriteString("[")
	for i, r := range repos {
		_, _ = sb.WriteString(r.GetFullName())
		if i+1 < len(repos) {
			_, _ = sb.WriteString(", ")
		}
	}
	_, _ = sb.WriteString("]")
	println("::set-output name=repositories::" + sb.String())
}
