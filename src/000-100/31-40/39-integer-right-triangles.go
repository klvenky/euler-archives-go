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
		for a := 1; a <= perimeter; a++ {
			bMax := perimeter - a
			for b := a; b <= bMax; b++ {
				cMax := bMax - b
				for c := b; c <= cMax; c++ {
					if isRightTriangle(a, b, c, perimeter) {
						tmp := []int{a, b, c}
						sort.Ints(tmp)
						eq := 0
						push := false
						if len(points) == 0 {
							push = true
						} else {
							for index, _ := range points {
								for tmpi, _ := range tmp {
									if tmp[tmpi] == points[index][tmpi] {
										eq++
									}
								}
								push = eq == 0
							}
						}
						if push {
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
	}
	fmt.Println("max perimter ---- ", currMaxPerimeter, "\nmax solutions", maxSolutions)

}

func isRightTriangle(a int, b int, c int, p int) bool {
	return (a+b+c == p) && (a*a+b*b) == c*c
}
