func numSpecial(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	rows_1s, cols_1s := make([]int, m), make([]int, n) // rows_1s[i] = k rows[i] has k ones
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				rows_1s[i]++
				cols_1s[j]++
			}
		}
	}
	special_cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && rows_1s[i] == 1 && cols_1s[j] == 1 {
				special_cnt++
			}
		}
	}
	return special_cnt
}