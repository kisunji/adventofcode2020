package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	// "cid",
}

var validators = map[string]func(string) bool{
	"byr": func(s string) bool {
		byr, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if !(byr >= 1920 && byr <= 2002) {
			return false
		}
		return true
	},
	"iyr": func(s string) bool {
		iyr, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if !(iyr >= 2010 && iyr <= 2020) {
			return false
		}
		return true
	},
	"eyr": func(s string) bool {
		eyr, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if !(eyr >= 2020 && eyr <= 2030) {
			return false
		}
		return true
	},
	"hgt": func(s string) bool {
		if strings.HasSuffix(s, "cm") {
			hgt, err := strconv.Atoi(strings.TrimSuffix(s, "cm"))
			if err != nil {
				return false
			}
			if !(hgt >= 150 && hgt <= 193) {
				return false
			}
			return true
		}
		if strings.HasSuffix(s, "in") {
			hgt, err := strconv.Atoi(strings.TrimSuffix(s, "in"))
			if err != nil {
				return false
			}
			if !(hgt >= 59 && hgt <= 76) {
				return false
			}
			return true
		}
		return false
	},
	"hcl": func(s string) bool {
		valid, err := regexp.MatchString("^#([a-f0-9]{6})$", s)
		if err != nil {
			return false
		}
		if !valid {
			return false
		}
		return true
	},
	"ecl": func(s string) bool {
		for _, valid := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if s == valid {
				return true
			}
		}
		return false
	},
	"pid": func(s string) bool {
		valid, err := regexp.MatchString("^([0-9]{9})$", s)
		if err != nil {
			return false
		}
		if !valid {
			return false
		}
		return true
	},
}

func main() {
	passports := parseInput()
	var validPassports int
	for _, p := range passports {
		isValid := true
		for _, f := range requiredFields {
			if val, ok := p[f]; !ok {
				isValid = false
				break
			} else {
				if !validators[f](val) {
					fmt.Printf("field %s:%s failed validation\n", f, val)
					isValid = false
					break
				}
			}
		}
		if isValid {
			validPassports++
		}
	}
	fmt.Println(validPassports)
}

func main_1() {
	passports := parseInput()
	var validPassports int
	for _, p := range passports {
		isValid := true
		for _, f := range requiredFields {
			if _, ok := p[f]; !ok {
				isValid = false
				break
			}
		}
		if isValid {
			validPassports++
		}
	}
	fmt.Println(validPassports)
}

type passport map[string]string

func parseInput() []passport {
	pps := strings.Split(input, "\n\n")
	var passports []passport
	for _, pp := range pps {
		passport := map[string]string{}
		kvs := strings.Fields(pp)
		for _, kv := range kvs {
			ss := strings.Split(kv, ":")
			passport[ss[0]] = ss[1]
		}
		passports = append(passports, passport)
	}
	return passports
}
