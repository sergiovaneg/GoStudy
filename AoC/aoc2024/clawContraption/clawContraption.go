package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

const EPS = 1e-7
const corrFactor = 1e13

func parseMachine(machine [3]string) (a mat.Matrix, b mat.Vector) {
	dRe := regexp.MustCompile(`\d+`)
	var nums [6]float64

	for idx, line := range machine {
		xy := dRe.FindAllString(line, 2)
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])

		nums[idx<<1] = float64(x)
		nums[idx<<1+1] = float64(y)
	}

	var aData [16]float64
	for idx, num := range nums[:4] {
		aData[(idx>>1)<<2+2+idx&0x01] = num
	}

	a = mat.NewSymDense(4, aData[:])
	b = mat.NewVecDense(4, []float64{-3, -1, nums[4], nums[5]})

	return
}

func getIntegerSolution(a mat.Matrix, x, b mat.Vector) int {
	m, n := math.Round(x.AtVec(0)), math.Round(x.AtVec(1))

	bRound := mat.NewVecDense(4, nil)
	bRound.SetVec(0, m)
	bRound.SetVec(1, n)

	bRound.MulVec(a, bRound)
	bRound.SubVec(bRound, b)

	if math.Abs(bRound.AtVec(2)) > EPS {
		return 0
	}
	if math.Abs(bRound.AtVec(3)) > EPS {
		return 0
	}

	return 3*int(m) + int(n)
}

func solveMinTokens(a mat.Matrix, b mat.Vector, correctionFlag bool) int {
	x := mat.NewVecDense(4, nil)

	if correctionFlag {
		aux := mat.NewVecDense(4, []float64{0, 0, corrFactor, corrFactor})
		aux.AddVec(b, aux)
		b = aux
	}

	if err := x.SolveVec(a, b); err != nil {
		fmt.Println(err)
		return 0
	}

	return getIntegerSolution(a, x, b)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	machine, r0, r1 := [3]string{}, 0, 0
	for idx := 0; scanner.Scan(); idx++ {
		if scanner.Text() == "" {
			a, b := parseMachine(machine)
			r0 += solveMinTokens(a, b, false)
			r1 += solveMinTokens(a, b, true)
			idx = -1
		} else {
			machine[idx] = scanner.Text()
		}
	}

	a, b := parseMachine(machine)
	r0 += solveMinTokens(a, b, false)
	r1 += solveMinTokens(a, b, true)

	println(r0)
	println(r1)
}
