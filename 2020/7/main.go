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

type bag struct {
	color    string
	contains map[*bag]int
}

func findFather(bags []string, color string) []string {
	fathers := make([]string, 0)

	for _, b := range bags {
		if strings.Contains(b, color) && !strings.HasPrefix(b, color) {
			fathers = append(fathers, strings.TrimSpace(b[:strings.Index(b, "bags")]))
		}
	}

	return fathers
}

func findSons(rule string) map[string]int {
	if strings.Contains(rule, "no other bags") {
		return map[string]int{}
	}
	resp := make(map[string]int)

	rules := strings.Split(rule, "contain")

	bags := strings.Split(rules[1], ",")

	for _, r := range bags {
		r = strings.TrimSpace(r)
		sNumber := strings.Split(r, " ")
		number, err := strconv.Atoi(sNumber[0])
		if err != nil {
			panic(err)
		}
		color := r[len(sNumber[0]):strings.Index(r, "bag")]

		resp[strings.TrimSpace(color)] = number
	}

	return resp
}

func firstproblem(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	goldColor := "shiny gold"

	containGold := make([]string, 0)
	for _, l := range lines {
		if strings.Contains(l, goldColor) {

			bagColor := strings.TrimSpace(l[:strings.Index(l, "bags")])
			if bagColor != goldColor {
				containGold = append(containGold, bagColor)
			}
		}
	}

	fathers := make(map[string]bool)
	for {
		for _, c := range containGold {
			fathers[c] = true
		}

		newFathers := make([]string, 0)

		for _, c := range containGold {
			newFathers = append(newFathers, findFather(lines, c)...)
		}

		if len(newFathers) == 0 {
			break
		}
		containGold = newFathers
	}

	return len(fathers)
}

func findRule(lines []string, color string) string {
	for _, l := range lines {
		if strings.HasPrefix(l, color) {
			return l
		}
	}

	return ""
}

func secondproblem(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	goldRule := findRule(lines, "shiny gold")

	sons := findSons(goldRule)
	goldBag := bag{
		color:    "shiny gold",
		contains: make(map[*bag]int),
	}
	for k, v := range sons {
		newBag := bag{
			color:    k,
			contains: make(map[*bag]int),
		}
		goldBag.contains[&newBag] = v
	}

	fillBag(lines, &goldBag.contains)

	return calcBags(goldBag.contains)
}

func fillBag(lines []string, contains *map[*bag]int) {
	for k := range *contains {
		newSons := findSons(findRule(lines, k.color))
		for c, s := range newSons {
			newBag := bag{
				color:    c,
				contains: make(map[*bag]int),
			}
			k.contains[&newBag] = s
		}
		fillBag(lines, &k.contains)
	}
}

func calcBags(contains map[*bag]int) int {
	count := 0
	for k, v := range contains {
		if len(k.contains) == 0 {
			count += v
		} else {
			count += v + v*calcBags(k.contains)
		}
	}
	return count
}

func main() {
	path := filepath.Join(build.Default.GOPATH, "src", "github.com", "PFadel", "adventofcode", "2020", "7", "input")

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
