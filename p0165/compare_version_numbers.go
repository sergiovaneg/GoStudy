package p0165

import (
	"slices"
	"strconv"
)

func getEndNum(v []rune, idx int) ([2]int, error) {
	if idx == -1 {
		return [2]int{-1, 0}, nil
	}
	end := slices.Index(v[idx:], '.')
	if end == -1 {
		num, err := strconv.Atoi(string(v[idx:]))
		return [2]int{-1, num}, err
	}
	end += idx
	num, err := strconv.Atoi(string(v[idx:end]))
	return [2]int{end + 1, num}, err
}

func CompareVersion(version1 string, version2 string) int {
	v1, v2 := []rune(version1), []rune(version2)
	idx1, idx2 := 0, 0

	for idx1 != -1 || idx2 != -1 {
		res1, err1 := getEndNum(v1, idx1)
		if err1 != nil {
			return 0
		}

		res2, err2 := getEndNum(v2, idx2)
		if err2 != nil {
			return 0
		}

		if res1[1] < res2[1] {
			return -1
		}

		if res1[1] > res2[1] {
			return 1
		}
		idx1, idx2 = res1[0], res2[0]
	}
	return 0
}
