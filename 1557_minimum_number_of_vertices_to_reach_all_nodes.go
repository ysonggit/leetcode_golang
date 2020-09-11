func findSmallestSetOfVertices(n int, edges [][]int) []int {
	indegrees := make([]int, n)
	for _, edge := range edges {
		indegrees[edge[1]] = 1
	}
	smallest_set := []int{}
	for i := 0; i < n; i++ {
		if indegrees[i] == 0 {
			smallest_set = append(smallest_set, i)
		}
	}
	return smallest_set
}