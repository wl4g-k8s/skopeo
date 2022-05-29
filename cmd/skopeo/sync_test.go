package main

import (
	"fmt"
	"path"
	"testing"
)

func TestExtractScopedLevelPath(t *testing.T) {
	var scopedLevel = 2

	// Case1:
	var repoUri = "/docker.io/calico/node:v3.21.0/"
	fmt.Println("extractScopedLevePath :: " + ExtractScopedLevelPath(repoUri, scopedLevel))
	fmt.Println("vs path.Base :: " + path.Base(repoUri))

	// Case2:
	repoUri = "redis:7.0.0"
	fmt.Println("extractScopedLevePath :: " + ExtractScopedLevelPath(repoUri, scopedLevel))
}
