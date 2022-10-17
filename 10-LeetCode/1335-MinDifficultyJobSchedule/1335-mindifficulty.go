package mindifficulty

import (
	"fmt"
	"sort"
)

// You want to schedule a list of jobs in d days. Jobs are dependent (i.e To work on the
// ith job, you have to finish all the jobs j where 0 <= j < i).

// You have to finish at least one task every day. The difficulty of a job schedule is
// the sum of difficulties of each day of the d days. The difficulty of a day is the
// maximum difficulty of a job done on that day.

// You are given an integer array jobDifficulty and an integer d. The difficulty of the
// ith job is jobDifficulty[i].

// Return the minimum difficulty of a job schedule. If you cannot find a schedule for the jobs return -1.

// 1 <= jobDifficulty.length <= 300
// 0 <= jobDifficulty[i] <= 1000
// 1 <= d <= 10
func minDifficulty(jobDifficulty []int, d int) int {
	var difficultyCache = make(map[string]int)
	var solutions = make(map[string]int)
	return minDifficultyRecurse(jobDifficulty, d, solutions, difficultyCache)
}

func minDifficultyRecurse(jobs []int, days int, solutions map[string]int, difficultyCache map[string]int) int {

	RecursiveCalls++

	// base cases
	// 	days == 0 => return -1?
	//	days == 1 => all remaining jobs
	//
	//  days == 0; jobs !empty => -1?
	//  days == 0; jobs empty => 0
	//  days == 1; jobs !empty => take remaining jobs
	//  days == 1; jobs empty => -1

	// heuristic: len(jobs) < days => -1 (more days left than jobs remaining

	if len(jobs) < days {
		// more days than jobs remaining
		return -1
	} else if days == 0 {
		if len(jobs) > 0 {
			return -1
		} else {
			return 0
		}
	} else if days == 1 {
		// last day - take the remaining jobs
		return max(jobs, difficultyCache)
	} else {
		RecursiveNonBase++

		// check to see if we already have a solution for this (need to use both the
		// array & number of days as the key, since the answer can change based on
		// both parameters)
		var key string = ""
		if SortKey {
			// sorting the key doesn't actually seem to help, since most scenarios
			// all have unique values - slows it down since it copies the array and
			// then sorts it
			var temp = make([]int, len(jobs))
			copy(temp, jobs)
			sort.Ints(temp)
			key = fmt.Sprintf("%v-%d", temp, days)
		} else {
			key = fmt.Sprintf("%v-%d", jobs, days)
		}

		var cached, found = solutions[key]
		if found {
			RecursiveHits++
			return cached
		}
		RecursiveMisses++

		var bestCase int = -2

		// try taking each job in sequence (start at 1 since there needs
		// to be at least one job taken today)
		for i := 1; i < len(jobs); i++ {
			var todaysJobs = jobs[:i]
			var remaining = jobs[i:]

			var todaysDifficulty = max(todaysJobs, difficultyCache)

			// recurse and figure out the remainder of the jobs
			var remainingDifficulty = minDifficultyRecurse(remaining, days-1, solutions, difficultyCache)
			if remainingDifficulty != -1 {
				// only consider the scenario if it's possible to break it down

				var possibleSolutionDifficulty = todaysDifficulty + remainingDifficulty

				if bestCase == -2 {
					// no solution has been identified yet - by default, this
					// is the best solution
					bestCase = possibleSolutionDifficulty
				} else {
					// check to see if the next potential solution is better
					// than what was already guessed
					if bestCase > possibleSolutionDifficulty {
						bestCase = possibleSolutionDifficulty
					}
				}
			}
		}

		// if no solutions were found, return -1
		if bestCase == -2 {
			bestCase = -1
		}
		solutions[key] = bestCase // cache this for the future
		return bestCase
	}
}

// caching the results of this call actually slows this down quite a bit (at least
// 2x) - probably because creating the key ends up iterating the collection once
// and then needing to do it again to calculate the result
func max(jobs []int, cache map[string]int) int {
	MaxCalls++

	if len(jobs) > 0 {
		var max = jobs[0]
		for _, job := range jobs[1:] {
			if job > max {
				max = job
			}
		}
		return max
	}
	return 0
}

var MaxCalls int = 0

var RecursiveCalls int = 0
var RecursiveNonBase int = 0
var RecursiveMisses int = 0
var RecursiveHits int = 0

var SortKey bool = false
