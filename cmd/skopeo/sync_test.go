package main

import (
	"fmt"
	"path"
	"testing"
)

func TestExtractScopedLevelPath(t *testing.T) {
	var scopedLevel = 2
	// var repoUri = "/registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:v1.18.1/"
	var repoUri = "/docker.io/calico/node:v3.21.0/"

	fmt.Println("extractScopedLevePath :: " + ExtractScopedLevelPath(repoUri, scopedLevel))

	fmt.Println("vs path.Base :: " + path.Base(repoUri))
}
