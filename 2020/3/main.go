package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

func toboggan(input string, stepX, stepY int) int {
	trees := 0
	x := 0
	y := 0

	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	linesSize := len(lines[0])

	for y < len(lines) {
		if fmt.Sprintf("%c", lines[y][x]) == "#" {
			trees++
		}

		y = y + stepY
		x = (x + stepX) % linesSize
	}
	return trees
}

func firstproblem(input string) int {
	return toboggan(input, 3, 1)
}

func secondproblem(input string) int {
	return toboggan(input, 1, 1) * toboggan(input, 3, 1) * toboggan(input, 5, 1) * toboggan(input, 7, 1) * toboggan(input, 1, 2)
}

func main() {
	path := filepath.Join(build.Default.GOPATH, "src", "github.com", "PFadel", "adventofcode", "2020", "3", "input")

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

	fmt.Printf("%d, %d Nanoseconds\n", r1, elapsed1.Nanoseconds())
	fmt.Printf("%d, %d Nanoseconds\n", r2, elapsed2.Nanoseconds())
}
