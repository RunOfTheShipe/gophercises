package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum_ScenarioOne(t *testing.T) {
	RunTest([]int{2, 7, 11, 15}, 9, []int{0, 1}, t)
}

func TestTwoSum_ScenarioTwo(t *testing.T) {
	RunTest([]int{3, 2, 4}, 6, []int{1, 2}, t)
}

func TestTwoSum_ScenarioThree(t *testing.T) {
	RunTest([]int{3, 3}, 6, []int{0, 1}, t)
}

func RunTest(nums []int, target int, expected []int, t *testing.T) {
	var actual = twoSum(nums, target)

	assert.ElementsMatchf(t, expected, actual, "Different contents!")
}
