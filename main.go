package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/exec"
	"runtime"

	crypto_rand "crypto/rand"
	math_rand "math/rand"
)

var (
	minQuestion = flag.Int("min_question", 1, "The lowest question")
	maxQuestion = flag.Int("max_question", 100, "The highest question")
	cmd         *exec.Cmd
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

	randomBytes := make([]byte, 1024)
	crypto_rand.Read(randomBytes)
	seed, err := crypto_rand.Int(bytes.NewBuffer(randomBytes), big.NewInt(1E18))
	if err != nil {
		log.Fatal(err)
	}

	r := math_rand.New(math_rand.NewSource(seed.Int64()))
	indexes := r.Perm(100)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("To jump between questions and answers, simply press Enter.")
	clearScreen()
	reader.ReadString('\n')

	for _, index := range indexes {
		if index == 0 || index < *minQuestion || index > *maxQuestion {
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

			reader.ReadString('\n')
		}
	}
}
