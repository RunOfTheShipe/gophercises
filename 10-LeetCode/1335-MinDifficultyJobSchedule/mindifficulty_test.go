package mindifficulty

import (
	"fmt"
	"testing"
)

func Test_S001(t *testing.T) {
	// First day you can finish the first 5 jobs, total difficulty = 6.
	// Second day you can finish the last job, total difficulty = 1.
	// The difficulty of the schedule = 6 + 1 = 7
	compareMinDifficultyWithSorting(t, []int{6, 5, 4, 3, 2, 1}, 2, 7)
}

func Test_S002(t *testing.T) {
	// If you finish a job per day you will still have a free day. you
	// cannot find a schedule for the given jobs.
	compareMinDifficultyWithSorting(t, []int{9, 9, 9}, 4, -1)
}

func Test_S003(t *testing.T) {
	// The schedule is one job per day. total difficulty will be 3.
	compareMinDifficultyWithSorting(t, []int{1, 1, 1}, 3, 3)
}

func Test_S004(t *testing.T) {
	compareMinDifficultyWithSorting(t, []int{0, 0, 0}, 3, 0)
}

func Test_T017(t *testing.T) {
	compareMinDifficultyWithSorting(t, []int{11, 111, 22, 222, 33, 333, 44, 444}, 6, 843)
}

func Test_T025(t *testing.T) {

	var jobs = []int{749, 811, 666, 27, 594, 696, 572, 886, 198, 761, 292, 542, 257, 470, 408, 145, 26, 677, 34, 577, 758, 175, 558, 247, 493, 274, 629, 436, 797, 108, 306, 111, 104, 435, 529, 158, 73, 24, 704, 646, 275, 230, 301, 467, 919, 446, 851, 868, 971, 195, 838, 799, 685, 428, 942, 813, 499, 756, 733, 508, 823, 884, 539, 137, 997, 533, 236, 92, 169, 708, 237, 427, 546, 380, 505, 902, 159, 944, 802, 520, 830, 550, 968, 951, 991, 40, 691, 212, 788, 748, 906, 839, 400, 367, 984, 454, 171, 785, 931, 841, 583, 876, 490, 112, 184, 97, 811, 83, 465, 889, 679, 203, 496, 434, 342, 591, 313, 940, 223, 534, 737, 182, 119, 331, 63, 130, 459, 913, 558, 418, 576, 918, 472, 437, 490, 805, 188, 252, 447, 256, 384, 464, 488, 760, 943, 587, 512, 316, 80, 513, 542, 846, 118, 56, 634, 601, 380, 74, 489, 667, 446, 518, 106, 949, 825, 126, 51, 626, 285, 8, 770, 302, 54, 981, 699, 765, 899, 996, 319, 675, 908, 93, 756, 151, 793, 942, 68, 927, 732, 353, 121, 795, 359, 569, 603, 195, 789, 531, 84, 613, 827, 647, 728, 307, 469, 449, 22, 161, 164, 114, 367, 703, 384, 979, 1000, 703, 380, 533, 984, 432, 906, 481, 92, 123, 293, 799, 63, 863, 860, 255, 842, 907, 293, 739, 316, 699, 236, 714, 475, 154}

	compareMinDifficultyWithSorting(t, jobs, 10, 3295)
}

// ----- helpers ----- //
func runMinDifficultyTest(t *testing.T, jobs []int, d int, expected int) {
	t.Helper()

	// reset our counters
	MaxCalls = 0

	RecursiveCalls = 0
	RecursiveNonBase = 0
	RecursiveMisses = 0
	RecursiveHits = 0

	var actual int = minDifficulty(jobs, d)

	t.Logf("Calls to Recurse(): Calls=%d (Non-Base=%d Hits=%d Misses=%d) max(): Calls=%d",
		RecursiveCalls, RecursiveNonBase, RecursiveHits, RecursiveMisses,
		MaxCalls)

	if actual != expected {
		t.Errorf("Jobs=%v d=%d Expected=%d Actual=%d", jobs, d, expected, actual)
	}
}

var sortKeyEnumeration = []struct {
	sortBeforeStart bool
}{
	{sortBeforeStart: true},
	{sortBeforeStart: false},
}

func compareMinDifficultyWithSorting(t *testing.T, jobs []int, days int, expected int) {
	t.Helper()
	for _, value := range sortKeyEnumeration {
		t.Run(fmt.Sprintf("sort-key-%v", value.sortBeforeStart), func(t1 *testing.T) {
			SortKey = value.sortBeforeStart

			runMinDifficultyTest(t, jobs, days, expected)
		})
	}
}
