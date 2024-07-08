package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Room struct {
	id       int
	name     string
	checksum string
}

func parseRoom(line string) Room {
	idLoc := regexp.MustCompile("([0-9]+)").FindStringIndex(line)
	id, _ := strconv.Atoi(line[idLoc[0]:idLoc[1]])

	return Room{
		id:       id,
		name:     line[:idLoc[0]-1],
		checksum: line[idLoc[1]+1 : idLoc[1]+6],
	}
}

func (r Room) genChecksum() string {
	count := make(map[rune]int, 26)
	for _, char := range r.name {
		if char == '-' {
			continue
		}
		count[char]++
	}

	checksum := make([]rune, 0, 26)
	for char := rune('a'); char <= rune('z'); char++ {
		checksum = append(checksum, char)
	}

	slices.SortFunc(checksum, func(a, b rune) int {
		if count[a] == count[b] {
			return int(a) - int(b)
		}
		return count[b] - count[a]
	})

	return string(checksum[:5])
}

func (r Room) decryptName() string {
	shift := r.id % 26

	name := ""
	for _, char := range r.name {
		if char == '-' {
			name += " "
			continue
		}
		c := rune((int(char-'a')+shift)%26) + 'a'
		name += string(c)
	}

	return name
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	validRooms := make([]Room, 0)
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		r := parseRoom(scanner.Text())
		if r.checksum == r.genChecksum() {
			result += r.id
			validRooms = append(validRooms, r)
		}
	}

	println(result)

	for _, room := range validRooms {
		decryptedName := room.decryptName()
		if strings.Contains(decryptedName, "north") {
			fmt.Printf("Name: %v\tId: %v\n", decryptedName, room.id)
		}
	}
}
