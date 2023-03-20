package placeflowers

import "testing"

func Test001(t *testing.T) {
	runTest([]int{1, 0, 0, 0, 1}, 1, true, t)
}

func Test002(t *testing.T) {
	runTest([]int{1, 0, 0, 0, 1}, 2, false, t)
}

func Test014(t *testing.T) {
	runTest([]int{0}, 1, true, t)
}

func Test014a(t *testing.T) {
	runTest([]int{0, 0}, 1, true, t)
}

func Test014b(t *testing.T) {
	runTest([]int{}, 1, false, t)
}

func Test014c(t *testing.T) {
	runTest([]int{}, 0, true, t)
}

func Test105(t *testing.T) {
	runTest([]int{1}, 0, true, t)
}

func Test111(t *testing.T) {
	runTest([]int{0, 0, 1, 0, 0}, 1, true, t)
}

func runTest(flowerbed []int, n int, expected bool, t *testing.T) {
	t.Helper()

	var actual = canPlaceFlowers(flowerbed, n)
	if actual != expected {
		t.Errorf("Expected: %v Actual: %v", expected, actual)
	}
}
