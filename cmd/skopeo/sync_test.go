package main

import (
	"fmt"
	"path"
	"strings"
	"testing"
)

func TestExtractScopedLevelPath(t *testing.T) {
	var scopedLevel = 6

	// Case1:
	var repoUri = "/docker.io/calico/node:v3.21.0/"
	fmt.Println("extractScopedLevePath :: " + ExtractScopedLevelPath(repoUri, scopedLevel))
	fmt.Println("vs path.Base :: " + path.Base(repoUri))

	// Case2:
	repoUri = "redis:7.0.0"
	fmt.Println("extractScopedLevePath :: " + ExtractScopedLevelPath(repoUri, scopedLevel))

	// Case3:
	repoUri = "gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/controller"
	fmt.Println("extractScopedLevePath :: " + ExtractScopedLevelPath(repoUri, scopedLevel))
}

func TestSubstring(t *testing.T) {
	dockerRepoDefault := "docker.io/library/"
	destSuffix := "docker.io/library/redis:3.0-alpine"
	idx := strings.Index(destSuffix, dockerRepoDefault)
	fmt.Println(idx)
	fmt.Println(destSuffix[idx+len(dockerRepoDefault):])

	fmt.Println(path.Base(destSuffix))
}
