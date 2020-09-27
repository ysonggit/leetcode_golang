func findJudge(N int, trust [][]int) int {
	cnt := make([]int, N+1)
	for i := 0; i < len(trust); i++ {
		cnt[trust[i][1]]++ // indegree
		cnt[trust[i][0]]-- // outdegree
	}
	for i := 1; i <= N; i++ {
		if cnt[i] == N-1 {
			return i
		}
	}
	return -1
}