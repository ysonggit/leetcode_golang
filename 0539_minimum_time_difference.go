import (
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	min_diff := 24*60 + 1
	n := len(timePoints)
	for i := n - 1; i > 0; i-- {
		diff := timeInMin(timePoints[i]) - timeInMin(timePoints[i-1])
		if diff < min_diff {
			min_diff = diff
		}
	}
	overnight_diff := (timeInMin(timePoints[0]) + 24*60) - timeInMin(timePoints[n-1])
	if min_diff > overnight_diff {
		min_diff = overnight_diff
	}
	return min_diff
}

func timeInMin(time string) int {
	hh, _ := strconv.Atoi(time[0:2])
	mm, _ := strconv.Atoi(time[3:])
	return hh*60 + mm
} 