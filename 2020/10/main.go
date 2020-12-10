package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
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
	Jolt          int   `json:"jolt"`
	Possiblevolts []int `json:"possiblevolts"`
}

var adaptersMap map[int]adapter

func findAdapters(values valuesList, a *adapter) {
	for i, v := range values {
		if _, ok := adaptersMap[v]; ok {
			continue
		} else {
			if v-a.Jolt == 1 || v-a.Jolt == 2 || v-a.Jolt == 3 {
				new := adapter{
					Jolt:          v,
					Possiblevolts: []int{},
				}

				findAdapters(values[i:], &new)

				a.Possiblevolts = append(a.Possiblevolts, new.Jolt)
				if i+1 < len(values) && (values[i+1]-a.Jolt == 1 || values[i+1]-a.Jolt == 2 || values[i+1]-a.Jolt == 3) {
					new := adapter{
						Jolt:          values[i+1],
						Possiblevolts: []int{},
					}

					findAdapters(values[i+1:], &new)
					a.Possiblevolts = append(a.Possiblevolts, new.Jolt)
				}
				if i+2 < len(values) && (values[i+2]-a.Jolt == 1 || values[i+2]-a.Jolt == 2 || values[i+2]-a.Jolt == 3) {
					new := adapter{
						Jolt:          values[i+2],
						Possiblevolts: []int{},
					}

					findAdapters(values[i+2:], &new)
					a.Possiblevolts = append(a.Possiblevolts, new.Jolt)
				}
			}
			adaptersMap[a.Jolt] = *a
		}
	}
}

func ends(target int, a *adapter) int {
	if a.Jolt+3 == target {
		return 1
	}

	end := 0
	for _, s := range a.Possiblevolts {
		if s+3 == 1 {
			return 1
		}
		new := adaptersMap[s]
		end += ends(target, &new)
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
		Jolt:          0,
		Possiblevolts: []int{},
	}

	adaptersMap = make(map[int]adapter)
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
