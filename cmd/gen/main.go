package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("questions.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("package main")
	fmt.Println("var questions = map[int]map[string][]string{")

	questionRegex := regexp.MustCompile("^([0-9]*\\.)\\s+(.*)")
	answerRegex := regexp.MustCompile("^(▪)\\s+(.*)")

	count := 1
	for _, line := range strings.Split(string(content), "\n") {
		line = strings.TrimSpace(line)
		bytes := []byte(line)
		if ok := questionRegex.Match(bytes); ok {
			if count > 1 {
				fmt.Println("}},")
			}

			fmt.Printf("%d: {\"%s\": {\n", count, questionRegex.ReplaceAllString(line, "$2"))
			count++
		} else {
			fmt.Printf("\"%s\", \n", answerRegex.ReplaceAllString(line, "$2"))
		}

	}

	fmt.Printf("}},\n}")
}
