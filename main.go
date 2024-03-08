package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func getCurrentBranch() string {
	output := new(strings.Builder)

	c := exec.Command("git", "branch", "--show-current")
	c.Stdout = output
	c.Stderr = output
	if err := c.Run(); err != nil {
		log.Println(output)
		log.Fatal(err)
	}
	return strings.TrimRight(output.String(), "\n")

}

var commitRegexp = regexp.MustCompile(`((?P<change>\w*)\/)((?P<domain>.*)\/)?(?P<title>.*)`)

func replaceDashes(str string) string {
	re := regexp.MustCompile(`-{1,}`)
	replaced := re.ReplaceAllStringFunc(str, func(s string) string {
		if len(s) == 1 {
			return " "
		} else {
			return "-"
		}
	})
	return replaced
}

func parse(branchName string) string {
	matches := commitRegexp.FindStringSubmatch(branchName)
	result := make(map[string]string)
	for i, name := range commitRegexp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
	}
	change := result["change"]
	domain := replaceDashes(result["domain"])
	title := replaceDashes(result["title"])

	if domain != "" {
		return fmt.Sprintf("%s(%s): %s", change, domain, title)
	} else {
		return fmt.Sprintf("%s: %s", change, title)
	}
}

func main() {
	branchName := getCurrentBranch()
	fmt.Println(parse(branchName))
}
