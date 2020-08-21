func busyStudent(startTime []int, endTime []int, queryTime int) int {
	cnt := 0
	for i := 0; i < len(endTime); i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			cnt++
		}
	}
	return cnt
}