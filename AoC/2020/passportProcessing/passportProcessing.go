package main

import (
	"bufio"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type ValFn func(string, string) bool

func policyB(line, field string) bool {
	line += " "
	lb := strings.Index(line, field+":") + len(field) + 1
	ub := lb + strings.Index(line[lb:], " ")
	val := line[lb:ub]

	switch field {
	case "byr":
		if len(val) != 4 {
			return false
		}
		num, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return num >= 1920 && num <= 2002

	case "iyr":
		if len(val) != 4 {
			return false
		}
		num, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return num >= 2010 && num <= 2020

	case "eyr":
		if len(val) != 4 {
			return false
		}
		num, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		return num >= 2020 && num <= 2030

	case "hgt":
		if strings.HasSuffix(val, "cm") {
			lb, ub = 150, 193
		} else if strings.HasSuffix(val, "in") {
			lb, ub = 59, 76
		} else {
			return false
		}

		num, err := strconv.Atoi(val[:len(val)-2])
		if err != nil {
			return false
		}
		return num >= lb && num <= ub

	case "hcl":
		if len(val) != 7 {
			return false
		}

		return regexp.MustCompile(`\#[\da-f]{6}`).MatchString(val)

	case "ecl":
		return slices.Contains([]string{
			"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}, val)

	case "pid":
		return len(val) == 9 && regexp.MustCompile(`\d{9}`).MatchString(val)
	}

	return false
}

func hasValidFields(passport, fields []string, valfn ValFn) bool {
	for _, field := range fields {
		var flag bool
		for _, line := range passport {
			if !strings.Contains(line, field+":") {
				continue
			}

			if valfn(line, field) {
				flag = true
			}
			break
		}

		if !flag {
			return false
		}
	}

	return true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fields := []string{
		"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
	}

	passport := make([]string, 0)
	var resA, resB int
	for scanner.Scan() {
		if scanner.Text() == "" {
			if hasValidFields(
				passport, fields,
				func(_, _ string) bool { return true }) {
				resA++

				if hasValidFields(passport, fields, policyB) {
					resB++
				}
			}

			passport = make([]string, 0)
		} else {
			passport = append(passport, scanner.Text())
		}
	}

	println(resA)
	println(resB)
}
