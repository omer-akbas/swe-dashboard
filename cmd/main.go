package main

import (
	"fmt"
	"os"
	"swe-dashboard/internal/metrics/selfmerging"
	"swe-dashboard/internal/scm/gitlab"
)

func main() {
	baseURL := os.Getenv("SWE_DASHBOARD_GITLAB_BASEURL")
	token := os.Getenv("SWE_DASHBOARD_GITLAB_TOKEN")

	gitlab, err := gitlab.NewSCM(gitlab.GitlabBaseURL(baseURL), gitlab.GitlabToken(token))
	if err != nil {
		panic(err)
	}

	self := selfmerging.NewSelfMergingService(gitlab)
	selfmrs, err := self.GetSelfMergingUsers("merged", "all", 10)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(selfmrs); i++ {
		fmt.Printf("%s\t%f\r\n", selfmrs[i].Username, selfmrs[i].Count)
	}
}
