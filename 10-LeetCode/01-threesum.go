package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	var length int = len(nums)

	var answer [][]int

	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			for k := j; k < length; k++ {
				if i != j && i != k && j != k {
					if nums[i]+nums[j]+nums[k] == 0 {

						var pa = []int{nums[i], nums[j], nums[k]}
						sort.Ints(pa)

						// iterate the existing answers and check to make sure it doesn't already exist
						var unique bool = true
						for a := 0; a < len(answer); a++ {
							if answer[a][0] == pa[0] && answer[a][1] == pa[1] && answer[a][2] == pa[2] {
								unique = false
							}
						}

						if unique {
							answer = append(answer, pa)
						}
					}
				}
			}
		}
	}

	return answer
}
