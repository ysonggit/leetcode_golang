func isRobotBounded(instructions string) bool {
	// https://leetcode.com/problems/robot-bounded-in-circle/discuss/850437/Python-O(n)-solution-explained
	x, y, dx, dy := 0, 0, 0, 1
	for t := 0; t < 4; t++ {
		for _, i := range instructions {
			if i == 'G' {
				x, y = x+dx, y+dy
			} else if i == 'L' {
				dx, dy = -dy, dx
			} else {
				dx, dy = dy, -dx
			}
		}
	}
	return (x == 0 && y == 0)
}
