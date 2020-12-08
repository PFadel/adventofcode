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

func executeProgram(lines []string) (int, bool) {
	acc := 0
	executed := make(map[int]int)
	i := 0

	for {
		if i >= len(lines) {
			return acc, true
		}

		if val, ok := executed[i]; ok && val > 1 {
			return -1, false
		}

		if _, ok := executed[i]; ok {
			executed[i]++
		} else {
			executed[i] = 1
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

	for i := range lines {
		changed := false
		if strings.HasPrefix(lines[i], "jmp") {
			changed = true
			lines[i] = strings.Replace(lines[i], "jmp", "nop", 1)
		}

		if acc, ok := executeProgram(lines); ok {
			return acc
		}

		if strings.HasPrefix(lines[i], "nop") && changed {
			lines[i] = strings.Replace(lines[i], "nop", "jmp", 1)
		}
	}

	return -1
}

func main() {
	path := filepath.Join(build.Default.GOPATH, "src", "github.com", "PFadel", "adventofcode", "2020", "8", "input")

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
