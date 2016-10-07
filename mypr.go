package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	version string = "0.0.1"
)

func getEnvVar(varName string) (result string) {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == varName {
			return pair[1]
		}
	}
	return ""
}

func openUrlByBrowser(url string) (result int) {
	result = 0
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	default:
		fmt.Println("Your PC is not supported.")
	}

	return result
}

func showVersion() {
	fmt.Println("mypr version", version)
}

type Options struct {
	Open    []int  `short:"o" long:"open"    description:"Open list of  pull requests they are assigned to you on a web browser"`
	Version []bool `short:"v" long:"version" description:"Show version"`
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)

	parser.Name = "mypr"
	parser.Usage = "[OPTIONS]"

	_, err := parser.Parse()
	if err != nil {
		os.Exit(0)
	}

	if opts.Version != nil {
		showVersion()
		os.Exit(0)
	}

	githubUserName := getEnvVar("MYPR_GITHUB_USERNAME")
	if githubUserName == "" {
		fmt.Println("MYPR_GITHUB_USERNAME is not specified.")
		os.Exit(0)
	}

	githubOrg := getEnvVar("MYPR_GITHUB_ORG")
	// githubOrg can be empty

	prurl := "https://github.com/pulls?q=is:open+is:pr"
	if githubOrg != "" {
		prurl += "+user:" + githubOrg
	}

	prurl += "+assignee:" + githubUserName

	// TODO: get organization and username from env
	//if opts.Open != nil {
	openUrlByBrowser(prurl)
	os.Exit(0)
	//}

	// TODO: show in console in json format
}
