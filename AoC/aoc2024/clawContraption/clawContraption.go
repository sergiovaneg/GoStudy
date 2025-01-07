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
	m, n := int(math.Round(x.AtVec(0))), int(math.Round(x.AtVec(1)))
	x1, y1 := int(a.At(0, 2)), int(a.At(0, 3))
	x2, y2 := int(a.At(1, 2)), int(a.At(1, 3))
	xp, yp := int(b.AtVec(2)), int(b.AtVec(3))

	if x1*m+x2*n != xp {
		return 0
	}

	if y1*m+y2*n != yp {
		return 0
	}

	return 3*m + n
}

func solveMinTokens(a mat.Matrix, b mat.Vector) int {
	x := mat.NewVecDense(4, nil)

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

	machine, result := [3]string{}, 0
	for idx := 0; scanner.Scan(); idx++ {
		if scanner.Text() == "" {
			result += solveMinTokens(parseMachine(machine))
			idx = -1
		} else {
			machine[idx] = scanner.Text()
		}
	}
	result += solveMinTokens(parseMachine(machine))

	println(result)
}
