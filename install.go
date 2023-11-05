package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"

	"github.com/gocolly/colly"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func install(ctx *cli.Context) error {
	versions, err := fetchVersions()
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	sdkDir := getSDKDirPath(home)
	installedVers, err := listInstalledVersions(sdkDir)
	if err != nil {
		return err
	}
	// すでにインストール済みのバージョンは除外する
	versions = removeDuplicates(versions, installedVers)

	// バージョンを選択させる
	promptSelect := promptui.Select{
		Label: "Choose a version",
		Items: versions,
		Templates: &promptui.SelectTemplates{
			Help: "Install NEW Golang Version",
		},
	}
	_, result, err := promptSelect.Run()
	if err != nil {
		if err == promptui.ErrAbort || err == promptui.ErrInterrupt {
			fmt.Println("Aborted")
			return nil
		}
		return err
	}

	// install確認
	promptConfirm := promptui.Prompt{
		Label:     fmt.Sprintf("Install %s", result),
		IsConfirm: true,
	}
	_, err = promptConfirm.Run()
	if err != nil {
		if err == promptui.ErrAbort || err == promptui.ErrInterrupt {
			fmt.Println("Aborted")
			return nil
		}
		return err
	}

	// install実行
	cmd := exec.Command("go", "install", fmt.Sprintf("golang.org/dl/go%s@latest", result))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	cmd = exec.Command(fmt.Sprintf("go%s", result), "download")
	output, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

// https://go.dev/dl/をスクレイピングして現在のOS/ARCに一致するバージョン一覧を取得する
func fetchVersions() ([]string, error) {
	currentOS := runtime.GOOS
	currentArch := runtime.GOARCH
	ext := "tar.gz"
	if currentOS == "windows" {
		ext = "zip"
	}
	suffix := fmt.Sprintf(".%s-%s.%s", currentOS, currentArch, ext)
	c := colly.NewCollector()
	var versions []string
	c.OnHTML("table tr", func(e *colly.HTMLElement) {
		filename := e.ChildText("td:first-child")
		if !strings.HasSuffix(filename, suffix) {
			return
		}
		filename = strings.TrimPrefix(strings.TrimSuffix(filename, suffix), "go")
		versions = append(versions, filename)
	})
	err := c.Visit("https://go.dev/dl/")
	if err != nil {
		return nil, err
	}
	sort.Sort(sort.Reverse(sort.StringSlice(versions)))
	return versions, nil
}

// s2に含まれる要素をs1から除外する
func removeDuplicates(s1, s2 []string) []string {
	m := make(map[string]bool)

	for _, v := range s2 {
		m[v] = true
	}

	result := []string{}
	for _, v := range s1 {
		if _, exists := m[v]; !exists {
			result = append(result, v)
		}
	}

	return result
}
