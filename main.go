package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	minQuestion  = flag.Int("min_question", 1, "The lowest question")
	maxQuestion  = flag.Int("max_question", 100, "The highest question")
	questionList = flag.String("question_list", "", "The list of the questions to display comma separated")
	cmd          *exec.Cmd
)

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	flag.Parse()

	indexes := rand.Perm(100)
	reader := bufio.NewReader(os.Stdin)
	yes := make([]int, 0, 100)
	no := make([]int, 0, 100)

	clearScreen()
	fmt.Println("To show the answers press Enter. If the answer was correct press 'y' otherwise press 'n'")
	fmt.Println("Press enter to continue")
	reader.ReadString('\n')

	indexesSet := make(map[int]bool)
	if *questionList != "" {
		strList := strings.Split(*questionList, ",")
		for _, s := range strList {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatalf("error decoding the question list: %s", err)
			}
			indexesSet[int(i)] = true
		}
	}

	for _, index := range indexes {
		if index == 0 || index < *minQuestion || index > *maxQuestion {
			continue
		}

		if *questionList != "" && !indexesSet[index] {
			continue
		}

		for q, answers := range questions[index] {
			clearScreen()
			fmt.Printf("%d. %s", index, q)

			reader.ReadString('\n')

			sep := "  "
			if len(answers) > 1 {
				sep = " -"
			}
			for _, answer := range answers {
				fmt.Printf("%s %s\n", sep, answer)
			}

			for {
				fmt.Print("\n\nWas the answer correct? [y/n] ")
				r, _, err := reader.ReadRune()
				if err != nil {
					log.Fatalf("error reading the answer: %s", err)
				}
				switch r {
				case 'y':
					yes = append(yes, index)
					goto done
				case 'n':
					no = append(no, index)
					goto done
				}

				clearScreen()
				fmt.Printf("%d. %s\n", index, q)

				sep := "  "
				if len(answers) > 1 {
					sep = " -"
				}
				for _, answer := range answers {
					fmt.Printf("%s %s\n", sep, answer)
				}
			}
		done:
			reader.Reset(os.Stdin)
		}

		fmt.Println("Here's the list of the answers that you did not answer correctly:")
		for _, index := range no {
			for q := range questions[index] {
				fmt.Printf("%d. %s\n", index, q)
			}
		}
	}
}
