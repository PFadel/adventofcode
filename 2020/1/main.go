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
	values := make([]int, 0)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line != "" {
			value, _ := strconv.Atoi(line)

			values = append(values, value)
		}
	}

	for _, v := range values {
		for _, v2 := range values {
			if v+v2 == 2020 {
				return v * v2
			}
		}
	}

	return -1
}

func secondproblem(input string) int {
	values := make([]int, 0)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line != "" {
			value, _ := strconv.Atoi(line)

			values = append(values, value)
		}
	}

	for _, v := range values {
		for _, v2 := range values {
			for _, v3 := range values {
				if v+v2+v3 == 2020 {
					return v * v2 * v3
				}
			}
		}
	}
	return -1
}

func main() {
	path := filepath.Join(build.Default.GOPATH, "src", "github.com", "PFadel", "adventofcode", "2020", "1", "input")

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
