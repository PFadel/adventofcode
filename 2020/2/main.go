package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func firstproblem(input string) int {
	valid := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line != "" {
			count := 0
			info := strings.Split(line, " ")
			numbers := strings.Split(info[0], "-")

			min, _ := strconv.Atoi(numbers[0])
			max, _ := strconv.Atoi(numbers[1])

			letter := info[1][:len(":")]

			for _, char := range info[2] {
				if fmt.Sprintf("%c", char) == letter {
					count++
				}
			}
			if count <= max && count >= min {
				valid++
			}
		}
	}

	return valid
}

func secondproblem(input string) int {
	valid := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line != "" {
			info := strings.Split(line, " ")
			numbers := strings.Split(info[0], "-")

			first, _ := strconv.Atoi(numbers[0])
			second, _ := strconv.Atoi(numbers[1])

			letter := info[1][:len(":")]

			if (fmt.Sprintf("%c", info[2][first-1]) == letter) != (fmt.Sprintf("%c", info[2][second-1]) == letter) {
				valid++
			}
		}
	}

	return valid
}

func main() {
	path := filepath.Join(build.Default.GOPATH, "src", "github.com", "PFadel", "adventofcode", "2020", "2", "input")

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
