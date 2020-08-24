func transpose(A [][]int) [][]int {
	m := len(A)
	n := len(A[0])
	T := make([][]int, n)
	for i := 0; i < n; i++ {
		T[i] = make([]int, m)
		for j := 0; j < m; j++ {
			T[i][j] = A[j][i]
		}
	}
	return T
}