package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func firstproblem(input string) int {
	valid := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line != "" {
			count := 0
			info := strings.Split(line, " ")
			numbers := strings.Split(info[0], "-")

			min, err := strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}
			max, err := strconv.Atoi(numbers[1])
			if err != nil {
				panic(err)
			}

			letter := info[1][:len(":")]

			for _, char := range info[2] {
				s := fmt.Sprintf("%c", char)
				if s == letter {
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

			first, err := strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}
			second, err := strconv.Atoi(numbers[1])
			if err != nil {
				panic(err)
			}

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

	println(firstproblem(payload))
	println(secondproblem(payload))
}
