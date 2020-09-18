func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	for i := 0; i < len(connections); i++ {
		union(parent, connections[i][0], connections[i][1])
	}
	orphans := 0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			orphans++
		}
	}
	return orphans - 1
}

func find(parent []int, cur int) int {
	for cur != parent[cur] {
		cur = parent[cur]
	}
	return cur
}

func union(parent []int, x int, y int) {
	root_x, root_y := find(parent, x), find(parent, y)
	if root_x == root_y {
		return
	}
	parent[root_x] = root_y
}