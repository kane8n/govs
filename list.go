package main

import (
	"os"
	"strings"
)

func listInstalledVersions(sdkDir string) ([]string, error) {
	files, err := os.ReadDir(sdkDir)
	if err != nil {
		return nil, err
	}

	versions := []string{}
	for _, f := range files {
		if f.IsDir() && strings.HasPrefix(f.Name(), "go") {
			versions = append(versions, strings.TrimPrefix(f.Name(), "go"))
		}
	}
	return versions, nil
}
