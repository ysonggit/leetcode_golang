func maximalNetworkRank(n int, roads [][]int) int {
	if len(roads) < 2 {
		return len(roads)
	}
	edges := make(map[int]map[int]bool)
	for _, road := range roads {
		a, b := road[0], road[1]
		if _, ok := edges[a]; !ok {
			edges[a] = make(map[int]bool)
		}
		if _, ok := edges[b]; !ok {
			edges[b] = make(map[int]bool)
		}
		edges[a][b] = true
		edges[b][a] = true
	}
	max_rank := 0
	for i := 0; i < n; i++ {
		rank_i := len(edges[i])
		for j := i + 1; j < n; j++ {
			rank_j := len(edges[j])
			if _, ok := edges[i][j]; ok {
				rank_j--
			}
			if rank_i+rank_j > max_rank {
				max_rank = rank_i + rank_j
			}
		}
	}
	return max_rank
}