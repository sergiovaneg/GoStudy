package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/sergiovaneg/GoStudy/utils"
	"gonum.org/v1/gonum/mat"
)

const EPS = 1e-7
const correctionFactor = 1e13

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
		aux := mat.NewVecDense(4, []float64{0, 0, correctionFactor, correctionFactor})
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
	n, _ := utils.LineCounter(file)
	nMachines := (n + 1) >> 2

	machine := [3]string{}
	c0, c1 := make(chan int, nMachines), make(chan int, nMachines)
	for scanner.Scan() {
		for idx := 0; idx < 3; idx++ {
			machine[idx] = scanner.Text()
			scanner.Scan()
		}
		a, b := parseMachine(machine)
		func() {
			c0 <- solveMinTokens(a, b, false)
			c1 <- solveMinTokens(a, b, true)
		}()
	}

	var r0, r1 int
	for range nMachines {
		r0 += <-c0
		r1 += <-c1
	}

	println(r0)
	println(r1)
}
