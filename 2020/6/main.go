package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

func firstproblem(input string) int {
	lines := strings.Split(input, "\n")

	count := 0
	countGroup := 0
	answers := make(map[string]bool)

	for _, l := range lines {
		if l != "" {
			for _, c := range l {
				if _, ok := answers[fmt.Sprintf("%c", c)]; !ok {
					answers[fmt.Sprintf("%c", c)] = true
					countGroup++
				}
			}
		} else {
			count += countGroup
			countGroup = 0
			answers = make(map[string]bool)
		}
	}
	return count
}

func secondproblem(input string) int {
	lines := strings.Split(input, "\n")

	count := 0
	peoplesInGroup := 0
	answers := make(map[string]int)

	for _, l := range lines {
		if l != "" {
			peoplesInGroup++
			for _, c := range l {
				if _, ok := answers[fmt.Sprintf("%c", c)]; ok {
					answers[fmt.Sprintf("%c", c)]++
				} else {
					answers[fmt.Sprintf("%c", c)] = 1
				}
			}
		} else {
			for _, a := range answers {
				if a == peoplesInGroup {
					count++
				}
			}

			peoplesInGroup = 0
			answers = make(map[string]int)
		}
	}
	return count
}

func main() {
	path := filepath.Join(build.Default.GOPATH, "src", "github.com", "PFadel", "adventofcode", "2020", "6", "input")

	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	payload := string(b)

	start := time.Now()
	r1 := firstproblem(payload)
	elapsed1 := time.Since(start)

	start = time.Now()
	r2 := secondproblem(payload)
	elapsed2 := time.Since(start)

	fmt.Printf("%d, %f Seconds\n", r1, elapsed1.Seconds())
	fmt.Printf("%d, %f Seconds\n", r2, elapsed2.Seconds())
}
