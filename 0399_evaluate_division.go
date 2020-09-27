func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	n := len(queries)
	res := make([]float64, n)
	for i := 0; i < len(equations); i++ {
		eq := equations[i]
		if _, ok := graph[eq[0]]; !ok {
			graph[eq[0]] = make(map[string]float64)
		}
		graph[eq[0]][eq[1]] = values[i]
		if values[i] != 0.0 {
			if _, ok := graph[eq[1]]; !ok {
				graph[eq[1]] = make(map[string]float64)
			}
			graph[eq[1]][eq[0]] = 1.0 / values[i]
		}
	}
	for i := 0; i < n; i++ {
		visited := make(map[string]bool)
		res[i] = dfs(queries[i][0], queries[i][1], graph, visited)
	}
	return res
}

// https://leetcode.com/problems/evaluate-division/discuss/171649/1ms-DFS-with-Explanations
func dfs(start string, end string, graph map[string]map[string]float64, visited map[string]bool) float64 {
	if visited[start] {
		return -1
	}
	if _, ok := graph[start][end]; ok {
		return graph[start][end]
	}
	visited[start] = true
	for node, weight := range graph[start] {
		if _, ok := visited[node]; !ok {
			prod_weight := dfs(node, end, graph, visited)
			if prod_weight != -1.0 {
				return weight * prod_weight
			}
		}
	}
	return -1
}