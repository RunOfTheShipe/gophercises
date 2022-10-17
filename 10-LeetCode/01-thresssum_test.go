package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum_ScenarioOne(t *testing.T) {
	RunThreeSum([]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}, t)
}

func TestThreeSum_ScenarioTwo(t *testing.T) {
	RunThreeSum([]int{0, 1, 1}, nil, t)
}

func TestThreeSum_ScenarioThree(t *testing.T) {
	RunThreeSum([]int{0, 0, 0}, [][]int{{0, 0, 0}}, t)
}

func RunThreeSum(nums []int, expected [][]int, t *testing.T) {

	var actual = threeSum(nums)
	assert.ElementsMatch(t, expected, actual, "Mismatch!")
}
