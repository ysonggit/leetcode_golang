func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, c := range moves { // assume only RLUD
		if c == 'R' {
			x++
		} else if c == 'L' {
			x--
		} else if c == 'U' {
			y++
		} else {
			y--
		}
	}
	return x == 0 && y == 0
}
