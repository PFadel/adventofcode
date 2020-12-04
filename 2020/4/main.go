package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

var validColors = []string{
	"amb",
	"blu",
	"brn",
	"gry",
	"grn",
	"hzl",
	"oth",
}

var validateFunctions = map[string]func(string) bool{
	"byr": func(value string) bool {
		v, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return v >= 1920 && v <= 2002
	},
	"iyr": func(value string) bool {
		v, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return v >= 2010 && v <= 2020
	},
	"eyr": func(value string) bool {
		v, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return v >= 2020 && v <= 2030
	},
	"hgt": func(value string) bool {
		if strings.Index(value, "cm") != -1 {
			v, err := strconv.Atoi(strings.TrimSuffix(value, "cm"))
			if err != nil {
				return false
			}
			return v >= 150 && v <= 193
		}
		if strings.Index(value, "in") != -1 {
			v, err := strconv.Atoi(strings.TrimSuffix(value, "in"))
			if err != nil {
				return false
			}
			return v >= 59 && v <= 76
		}
		return false
	},
	"hcl": func(value string) bool {
		if strings.HasPrefix(value, "#") {
			value = strings.TrimPrefix(value, "#")
			matched, err := regexp.Match(`^[a-f|0-9]{6}$`, []byte(value))
			if err != nil {
				panic(err)
			}
			return matched
		}
		return false
	},
	"ecl": func(value string) bool {
		for _, c := range validColors {
			if value == c {
				return true
			}
		}
		return false
	},
	"pid": func(value string) bool {
		matched, err := regexp.Match(`^[0-9]{9}$`, []byte(value))
		if err != nil {
			panic(err)
		}
		return matched
	},
	"cid": func(_ string) bool {
		return true
	},
}

func firstproblem(input string) int {
	invalid := 0

	lines := strings.Split(input, "\n")
	passports := make([]string, 0)

	passport := ""
	for _, l := range lines {

		if l != "" {
			passport += " " + l
			continue
		}
		passports = append(passports, passport)

		passport = ""
	}

	for _, p := range passports {
		for _, f := range requiredFields {
			if strings.Index(p, f) == -1 {
				invalid++
				break
			}
		}
	}

	return len(passports) - invalid
}

func validateField(value, field string) bool {
	return validateFunctions[field](value)
}

func secondproblem(input string) int {
	invalid := 0

	lines := strings.Split(input, "\n")
	passports := make([]string, 0)

	passport := ""
	for _, l := range lines {

		if l != "" {
			passport += " " + l
			continue
		}
		passports = append(passports, passport)

		passport = ""
	}

	for _, p := range passports {
		valid := true

		for _, f := range requiredFields {
			if strings.Index(p, f) == -1 {
				invalid++
				valid = false
				break
			}
		}
		if valid {
			fields := strings.Split(p, " ")
			sort.Strings(fields)
			for _, f := range fields {
				if f != "" {
					values := strings.Split(f, ":")
					if len(values) != 2 {
						panic("deu ruim")
					}
					key := values[0]
					value := values[1]

					if !validateField(value, key) {
						valid = false
						invalid++
						break
					}
				}
			}
		}
	}
	return len(passports) - invalid
}

func main() {
	path := filepath.Join(build.Default.GOPATH, "src", "github.com", "PFadel", "adventofcode", "2020", "4", "input")

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
