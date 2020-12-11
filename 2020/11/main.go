package main

import (
	"fmt"
	"go/build"
	"hash/fnv"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

func checkAdjacent(x, y int, lines []string) int {
	occupied := 0

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if x+i < len(lines) && x+i >= 0 {
				if y+j < len(lines[x]) && y+j >= 0 && (i != 0 || j != 0) {
					if fmt.Sprintf("%c", lines[x+i][y+j]) == "#" {
						occupied++
					}
				}
			}
		}
	}

	return occupied
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func firstproblem(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	hashs := make(map[uint32]bool)

	c := 0
	for {
		old := hash(strings.Join(lines, ";"))
		if _, ok := hashs[old]; ok {
			c++
			println("loop! ", c)
			// break
		} else {
			hashs[old] = true
		}

		copy := make([]string, 0)
		for _, l := range lines {
			copy = append(copy, l)
		}

		for i := range copy {
			for j := range copy[i] {
				switch fmt.Sprintf("%c", copy[i][j]) {
				case "L":
					lines[i] = replaceAtIndex(lines[i], '#', j)
				case "#":
					if checkAdjacent(i, j, lines) >= 4 {
						lines[i] = replaceAtIndex(lines[i], 'L', j)
					}
				}
			}
		}

		new := hash(strings.Join(lines, ";"))
		if old == new {
			break
		}
	}

	occupied := 0
	for _, l := range lines {
		println(l)
		for _, c := range l {
			if fmt.Sprintf("%c", c) == "#" {
				occupied++
			}
		}
	}
	return occupied
}

// func secondproblem(input string) int {

// }

func main() {
	path := filepath.Join(build.Default.GOPATH, "src", "github.com", "PFadel", "adventofcode", "2020", "11", "input")

	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	payload := string(b)

	start := time.Now()
	r1 := firstproblem(payload)
	elapsed1 := time.Since(start)

	// start = time.Now()
	// r2 := secondproblem(payload)
	// elapsed2 := time.Since(start)

	fmt.Printf("%d, %f Seconds\n", r1, elapsed1.Seconds())
	// fmt.Printf("%d, %f Seconds\n", r2, elapsed2.Seconds())
}
