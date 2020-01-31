package main

import (
	"fmt"
	"os"
	"log"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)


func main() {
	workspace := drone.Workspace{}
	repo := drone.Repo{}
	build := drone.Build{}
	job := drone.Job{}

	plugin.Param("workspace", &workspace)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("job", &job)
	plugin.MustParse()

	log.Println(job.Status)
}
