func minTimeToVisitAllPoints(points [][]int) int {
	n := len(points)
	t := 0
	for i := 0; i < n-1; i++ {
		cur, nex := points[i], points[i+1]
		delta_x, delta_y := abs(nex[0]-cur[0]), abs(nex[1]-cur[1])
		if delta_x > delta_y {
			t += delta_x
		} else {
			t += delta_y
		}
	}
	return t
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}