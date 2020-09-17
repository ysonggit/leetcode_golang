import "strconv"

func robotSim(commands []int, obstacles [][]int) int {
	x, y, dx, dy := 0, 0, 0, 1
	obs_map := createObstaclesMap(obstacles)
	max_dist := 0
	for _, cmd := range commands {
		if 1 <= cmd && cmd <= 9 {
			for i := 0; i < cmd; i++ {
				key := coord(x+dx, y+dy)
				if _, ok := obs_map[key]; ok {
					break
				}
				x += dx
				y += dy
			}
		}
		if cmd == -1 {
			dx, dy = dy, -dx // cw
		}
		if cmd == -2 {
			dx, dy = -dy, dx // ccw
		}
		max_dist = max(max_dist, dist2(x, y))
	}
	return max_dist
}

func createObstaclesMap(obstacles [][]int) map[string]bool {
	obs_map := make(map[string]bool)
	for _, obs := range obstacles {
		key := coord(obs[0], obs[1])
		obs_map[key] = true
	}
	return obs_map
}

func coord(x, y int) string {
	return strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

func dist2(x, y int) int {
	return x*x + y*y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}