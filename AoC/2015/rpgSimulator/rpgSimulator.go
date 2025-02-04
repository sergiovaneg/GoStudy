package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Item struct {
	cost int
	dmg  int
	def  int
}

const BossHP = 104
const BossDMG = 8
const BossDEF = 1

func parseItem(line string) Item {
	values := [3]int{}
	for idx, match := range regexp.MustCompile(
		"([0-9]+)").FindAllString(line, 3) {
		val, _ := strconv.Atoi(match)
		values[idx] = val
	}
	return Item{
		cost: values[0],
		dmg:  values[1],
		def:  values[2],
	}
}

func simulateFight(hp, dmg, def int) bool {
	bossHP := BossHP
	dmg1, dmg2 := max(1, dmg-BossDEF), max(1, BossDMG-def)
	for hp > 0 && bossHP > 0 {
		bossHP -= dmg1
		hp -= dmg2
	}

	return bossHP <= 0
}

func minimizeCost(weapons, armours, rings []Item) int {
	minimumCost := math.MaxInt
	var accCost int
	for _, weapon := range weapons {
		accCost += weapon.cost

		if accCost < minimumCost && simulateFight(
			100, weapon.dmg, 0) {
			minimumCost = accCost
		}

		for idx, ringOne := range rings {
			accCost += ringOne.cost

			if accCost < minimumCost && simulateFight(
				100,
				weapon.dmg+ringOne.dmg,
				ringOne.def) {
				minimumCost = accCost
			}

			for _, ringTwo := range rings[idx+1:] {
				if accCost+ringTwo.cost < minimumCost && simulateFight(
					100,
					weapon.dmg+ringOne.dmg+ringTwo.dmg,
					ringOne.def+ringTwo.def) {
					minimumCost = accCost + ringTwo.cost
				}
			}

			accCost -= ringOne.cost
		}

		for _, armour := range armours {
			accCost += armour.cost

			if accCost < minimumCost && simulateFight(
				100, weapon.dmg, armour.def) {
				minimumCost = accCost
			}

			for idx, ringOne := range rings {
				accCost += ringOne.cost

				if accCost < minimumCost && simulateFight(
					100,
					weapon.dmg+ringOne.dmg,
					armour.def+ringOne.def) {
					minimumCost = accCost
				}

				for _, ringTwo := range rings[idx+1:] {
					if accCost+ringTwo.cost < minimumCost && simulateFight(
						100,
						weapon.dmg+ringOne.dmg+ringTwo.dmg,
						armour.def+ringOne.def+ringTwo.def) {
						minimumCost = accCost + ringTwo.cost
					}
				}

				accCost -= ringOne.cost
			}

			accCost -= armour.cost
		}

		accCost -= weapon.cost
	}

	return minimumCost
}

func maximizeCost(weapons, armours, rings []Item) int {
	maximumCost := 0
	var accCost int
	for _, weapon := range weapons {
		accCost += weapon.cost

		if accCost > maximumCost && !simulateFight(
			100, weapon.dmg, 0) {
			maximumCost = accCost
		}

		for idx, ringOne := range rings {
			accCost += ringOne.cost

			if accCost > maximumCost && !simulateFight(
				100,
				weapon.dmg+ringOne.dmg,
				ringOne.def) {
				maximumCost = accCost
			}

			for _, ringTwo := range rings[idx+1:] {
				if accCost+ringTwo.cost > maximumCost && !simulateFight(
					100,
					weapon.dmg+ringOne.dmg+ringTwo.dmg,
					ringOne.def+ringTwo.def) {
					maximumCost = accCost + ringTwo.cost
				}
			}

			accCost -= ringOne.cost
		}

		for _, armour := range armours {
			accCost += armour.cost

			if accCost > maximumCost && !simulateFight(
				100, weapon.dmg, armour.def) {
				maximumCost = accCost
			}

			for idx, ringOne := range rings {
				accCost += ringOne.cost

				if accCost > maximumCost && !simulateFight(
					100,
					weapon.dmg+ringOne.dmg,
					armour.def+ringOne.def) {
					maximumCost = accCost
				}

				for _, ringTwo := range rings[idx+1:] {
					if accCost+ringTwo.cost > maximumCost && !simulateFight(
						100,
						weapon.dmg+ringOne.dmg+ringTwo.dmg,
						armour.def+ringOne.def+ringTwo.def) {
						maximumCost = accCost + ringTwo.cost
					}
				}

				accCost -= ringOne.cost
			}

			accCost -= armour.cost
		}

		accCost -= weapon.cost
	}

	return maximumCost
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	weapons := make([]Item, 0, 5)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		weapons = append(weapons, parseItem(line))
	}

	scanner.Scan()
	armours := make([]Item, 0, 5)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		armours = append(armours, parseItem(line))
	}

	scanner.Scan()
	rings := make([]Item, 0, 6)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		startIdx := strings.Index(line, "+")
		rings = append(rings, parseItem(line[startIdx+2:]))
	}

	println(minimizeCost(weapons, armours, rings))
	println(maximizeCost(weapons, armours, rings))
}
