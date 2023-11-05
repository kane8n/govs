package main

import (
	"os"
	"path"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func switchVersion(ctx *cli.Context) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	if err := createGovsDir(home); err != nil {
		return err
	}

	sdkDir := getSDKDirPath(home)

	installedVers, err := listInstalledVersions(sdkDir)
	if err != nil {
		return err
	}

	prompt := promptui.Select{
		Label: "Choose a version",
		Items: installedVers,
		Templates: &promptui.SelectTemplates{
			Help: "Golang Version Switcher",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		return err
	}

	goroot := path.Join(sdkDir, "go"+result)

	if err := setGOROOT(home, goroot); err != nil {
		return err
	}
	if err := setPATH(home, goroot); err != nil {
		return err
	}

	return nil
}

func setGOROOT(home, goroot string) error {
	if err := os.WriteFile(path.Join(home, ".govs", "goroot"), []byte(goroot), 0644); err != nil {
		return err
	}
	return nil
}

func setPATH(home, goroot string) error {
	gorootBin := path.Join(goroot, "bin")
	pathEnv := os.Getenv("PATH")
	pathMap := map[string]bool{}
	for _, p := range strings.Split(pathEnv, ":") {
		if p == gorootBin {
			continue
		}
		pathMap[p] = true
	}
	keys := []string{}
	for k := range pathMap {
		keys = append(keys, k)
	}
	if err := os.WriteFile(path.Join(home, ".govs", "path"), []byte(gorootBin+":"+strings.Join(keys, ":")), 0644); err != nil {
		return err
	}
	return nil
}
