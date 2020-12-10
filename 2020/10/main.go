package main

import (
	// "encoding/json"
	"fmt"
	"go/build"
	"io/ioutil"
	// "os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type valuesList []int

func (v valuesList) Len() int {
	return len(v)
}

func (v valuesList) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v valuesList) Swap(i, j int) {
	swap := v[i]
	v[i] = v[j]
	v[j] = swap
}

func firstproblem(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	values := make(valuesList, 0)
	for _, l := range lines {
		v, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		values = append(values, v)
	}
	sort.Sort(values)

	diff1 := 1
	diff2 := 0
	diff3 := 1
	old := 0
	for _, v := range values {
		if old != 0 {
			switch v - old {
			case 1:
				diff1++
			case 2:
				diff2++
			case 3:
				diff3++
			}
		}
		old = v
	}

	return diff1 * diff3
}

type adapter struct {
	Jolt             int       `json:"jolt"`
	PossibleAdapters []adapter `json:"possibleAdapters"`
}

func findAdapters(values valuesList, a *adapter) {
	for _, v := range values {
		if v-a.Jolt == 1 || v-a.Jolt == 2 || v-a.Jolt == 3 {
			new := adapter{
				Jolt:             v,
				PossibleAdapters: []adapter{},
			}

			findAdapters(values, &new)
			// output, err := json.Marshal(new)
			// if err != nil {
			// 	panic(err)
			// }
			// ioutil.WriteFile(strconv.Itoa(new.Jolt), output, os.ModePerm)

			a.PossibleAdapters = append(a.PossibleAdapters, new)
		}
	}
}

func ends(target int, a *adapter) int {
	if a.Jolt+3 == target {
		return 1
	}

	end := 0
	for _, s := range a.PossibleAdapters {
		if s.Jolt+3 == 1 {
			return 1
		}
		end += ends(target, &s)
	}
	return end
}

func secondproblem(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	values := make(valuesList, 0)
	for _, l := range lines {
		v, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		values = append(values, v)
	}
	sort.Sort(values)

	start := adapter{
		Jolt:             0,
		PossibleAdapters: []adapter{},
	}

	findAdapters(values, &start)

	return ends(values[len(values)-1]+3, &start)
}

func main() {
	// path := filepath.Join(build.Default.GOPATH, "src", "github.com", "PFadel", "adventofcode", "2020", "10", "input")
	path := filepath.Join(build.Default.GOPATH, "src", "github.com", "PFadel", "adventofcode", "2020", "10", "test")

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
