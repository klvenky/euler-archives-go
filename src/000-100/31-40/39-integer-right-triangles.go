/*
Integer right triangles
Problem 39

If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.
{20,48,52}, {24,45,51}, {30,40,50}
For which value of p ≤ 1000, is the number of solutions limitimised?
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	maxSolutions := 0
	currMaxPerimeter := 0
	limit := 1000
	for perimeter := 1; perimeter <= limit; perimeter++ {
		var points [][]int
		// perimeter := 120
		for a := 1; a <= perimeter; a++ {
			bMax := perimeter - a
			for b := a; b <= bMax; b++ {
				cMax := bMax - b
				for c := b; c <= cMax; c++ {
					isRight := isRightTriangle(a, b, c, perimeter)
					if isRight {
						//fmt.Println("a: ", a, "| b: ", b, "| c: ", c, " | p: ", perimeter)
						//fmt.Printf("{ %d, %d, %d }\n", a, b, c)
						tmp := []int{a, b, c}
						//fmt.Println("---before sort --- ", tmp)
						sort.Ints(tmp)
						//fmt.Println("---after sort --- ", tmp)

						eq := 0
						push := false
						if len(points) == 0 {
							push = true
						} else {
							//fmt.Println("in else ", len(points), " tmp -> ", tmp)
							for index, _ := range points {

								for tmpi, _ := range tmp {
									if tmp[tmpi] == points[index][tmpi] {
										eq++
									}
								}

								if eq == 0 {
									push = true
								} else {
									push = false
								}
								//fmt.Println("index -> ", index, " --- equal ->  ", eq, " --- push --> ", push)
							}
						}
						if push {
							//fmt.Println(" ---- pushing ---- ", tmp)
							points = append(points, tmp)
						}
					}
				}
			}
		}

		if len(points) > maxSolutions {
			maxSolutions = len(points)
			currMaxPerimeter = perimeter
		}
		// fmt.Println("perimeter : ", perimeter, " | solutions : ", maxSolutions)

	}
	fmt.Println("max perimter ---- ", currMaxPerimeter, "\nmax solutions", maxSolutions)

}

func isRightTriangle(a int, b int, c int, p int) bool {
	return (a+b+c == p) && (a*a+b*b) == c*c
}
