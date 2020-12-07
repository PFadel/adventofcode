package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"math"
	"path/filepath"
	"strings"
	"time"
)

func firstproblem(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	maxID := 0
	for _, l := range lines {
		minRow := 0.0
		maxRow := 127.0
		minSeat := 0.0
		maxSeat := 7.0

		for _, c := range l {
			switch fmt.Sprintf("%c", c) {
			case "F":
				maxRow -= math.Ceil((maxRow - minRow) / 2.0)
			case "B":
				minRow += math.Ceil((maxRow - minRow) / 2.0)
			case "L":
				maxSeat -= math.Ceil((maxSeat - minSeat) / 2.0)
			case "R":
				minSeat += math.Ceil((maxSeat - minSeat) / 2.0)
			}

		}

		if id := (int(maxRow) * 8) + int(maxSeat); id > maxID {
			maxID = id
		}
	}

	return maxID
}

func secondproblem(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	airplane := [128][8]bool{}

	for _, l := range lines {
		minRow := 0.0
		maxRow := 127.0
		minSeat := 0.0
		maxSeat := 7.0

		for _, c := range l {
			switch fmt.Sprintf("%c", c) {
			case "F":
				maxRow -= math.Ceil((maxRow - minRow) / 2.0)
			case "B":
				minRow += math.Ceil((maxRow - minRow) / 2.0)
			case "L":
				maxSeat -= math.Ceil((maxSeat - minSeat) / 2.0)
			case "R":
				minSeat += math.Ceil((maxSeat - minSeat) / 2.0)
			}

		}

		airplane[int(maxRow)][int(maxSeat)] = true
	}

	for p, i := range airplane {
		for k, j := range i {
			if !j && (p > 5 && p < 119) {
				return p*8 + k
			}
		}
	}

	return -1
}

func main() {
	path := filepath.Join(build.Default.GOPATH, "src", "github.com", "PFadel", "adventofcode", "2020", "5", "input")

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
