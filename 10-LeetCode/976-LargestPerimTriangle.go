package leetcode

import (
	"fmt"
	"math"
)

func largestPerimeter(nums []int) int {
	return recurse(nums, 0, 0)
	// var longest int = 0
	// var length int = len(nums)

	// for i := 0; i < length; i++ {
	// 	for j := i + 1; j < length; j++ {
	// 		for k := j + 1; k < length; k++ {

	// 			if i != j && i != k && j != k {

	// 				var pa = CalculatePerimeter([3]int{nums[i], nums[j], nums[k]})
	// 				if pa > longest {
	// 					longest = pa
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	//return longest
}

func recurse(sides []int, currentMax int, stackCount int) int {
	fmt.Printf("Stack=%v Sides=%v\n", stackCount, len(sides))

	var sidesLen = len(sides)
	if sidesLen < 3 {
		return 0
	} else if sidesLen == 3 {
		var localMax int = CalculatePerimeter(sides[0], sides[1], sides[2])
		if localMax > currentMax {
			return localMax
		} else {
			return currentMax
		}
	} else {
		return recurse(sides[1:sidesLen], currentMax, stackCount+1)
	}
}

func CalculatePerimeter(x int, y int, z int) int {
	if Area(x, y, z) > 0 {
		return x + y + z
	}
	return 0
}

func Area(x int, y int, z int) float64 {
	// convert everything to floats to do the maths
	var a float64 = float64(x)
	var b float64 = float64(y)
	var c float64 = float64(z)

	// heron's formula
	var s float64 = (a + b + c) / 2
	var inner float64 = s * (s - a) * (s - b) * (s - c)

	// this seems scary
	return math.Sqrt(inner)
}
