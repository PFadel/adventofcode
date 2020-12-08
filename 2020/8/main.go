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

func in(l []int, v int) bool {
	for _, x := range l {
		if x == v {
			return true
		}
	}
	return false
}

func remove(l []int, v int) []int {
	for i, x := range l {
		if x == v {
			return append(l[:i], l[i+1:]...)
		}
	}

	return l
}

func firstproblem(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	acc := 0
	executed := make([]int, 0)

	i := 0
	for {
		if in(executed, i) {
			return acc
		}
		executed = append(executed, i)
		l := lines[i]
		cmd := strings.Split(l, " ")

		switch cmd[0] {
		case "acc":
			sum, err := strconv.Atoi(cmd[1])
			if err != nil {
				panic(err)
			}
			acc += sum
			i++
		case "jmp":
			sum, err := strconv.Atoi(cmd[1])
			if err != nil {
				panic(err)
			}
			i += sum
		case "nop":
			i++
		}
	}
}

func secondproblem(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	acc := 0
	executed := make([]int, 0)

	i := 0
	// oldI := 0
	changedCmd := ""
	changedI := 0
	tried := make([]int, 0)
	for {
		if i >= len(lines) {
			break
		}

		if in(executed, i) {
			if changedCmd != "" {
				lines[changedI] = changedCmd
			}

			if !in(tried, i) {
				changedCmd = lines[i]
				changedI = i
				lines[i] = "nop +1"
				tried = append(tried, i)
			}

			executed = make([]int, 0)
			i = 0
			acc = 0
		} else {
			executed = append(executed, i)
		}

		l := lines[i]
		cmd := strings.Split(l, " ")

		switch cmd[0] {
		case "acc":
			sum, err := strconv.Atoi(cmd[1])
			if err != nil {
				panic(err)
			}
			acc += sum
			// oldI = i
			i++
		case "jmp":
			sum, err := strconv.Atoi(cmd[1])
			if err != nil {
				panic(err)
			}
			// oldI = i
			i += sum
		case "nop":
			// oldI = i
			i++
		}
	}

	return acc
}

func main() {
	path := filepath.Join(build.Default.GOPATH, "src", "github.com", "PFadel", "adventofcode", "2020", "8", "input")

	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	payload := string(b)

	// start := time.Now()
	// r1 := firstproblem(payload)
	// elapsed1 := time.Since(start)

	start := time.Now()
	r2 := secondproblem(payload)
	elapsed2 := time.Since(start)

	// fmt.Printf("%d, %f Seconds\n", r1, elapsed1.Seconds())
	fmt.Printf("%d, %f Seconds\n", r2, elapsed2.Seconds())
}
