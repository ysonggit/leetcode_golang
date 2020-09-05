func diagonalSum(mat [][]int) int {
	sum, n := 0, len(mat)
	for i := 0; i < n; i++ {
		sum += mat[i][i]
		if i != n-1-i {
			sum += mat[i][n-1-i]
		}
	}
	return sum
}