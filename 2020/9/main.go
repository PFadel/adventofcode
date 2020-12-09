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

const preambleSize = 25

func shift(l []int) []int {
	for i := range l {
		if i == len(l)-1 {
			l[i] = 0
		} else {
			l[i] = l[i+1]
		}
	}
	return l
}

func validatePreamble(preamble []int, value int) bool {
	for _, x := range preamble {
		for _, y := range preamble {
			if x+y == value && x != y {
				return true
			}
		}
	}
	return false
}

func firstproblem(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	values := make([]int, 0)
	for _, l := range lines {
		value, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		values = append(values, value)
	}

	preamble := make([]int, 0)
	preamble = append(preamble, values[:preambleSize]...)

	for _, v := range values[preambleSize:] {
		if !validatePreamble(preamble, v) {
			return v
		}
		preamble = shift(preamble)
		preamble[len(preamble)-1] = v
	}

	return 0
}

func secondproblem(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	values := make([]int, 0)
	for _, l := range lines {
		value, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		values = append(values, value)
	}

	invalidNumber := firstproblem(input)

	for j := range values {
		for i := j; i < len(values); i++ {
			sequence := make([]int, 0)
			sequence = append(sequence, values[j:i+1]...)
			sum := 0
			smallest := sequence[0]
			largest := 0

			for _, s := range sequence {
				sum += s
				if s > largest {
					largest = s
				}
				if s < smallest {
					smallest = s
				}
			}

			if sum == invalidNumber {
				return smallest + largest
			}
		}
	}

	return 0
}

func main() {
	path := filepath.Join(build.Default.GOPATH, "src", "github.com", "PFadel", "adventofcode", "2020", "9", "input")

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
