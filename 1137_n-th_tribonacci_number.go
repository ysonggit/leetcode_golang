func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n > 0 && n < 3 {
		return 1
	}
	t_i_3, t_i_2, t_i_1, t_i := 0, 1, 1, 0
	for i := 3; i <= n; i++ {
		t_i = t_i_1 + t_i_2 + t_i_3
		t_i_3, t_i_2, t_i_1 = t_i_2, t_i_1, t_i
	}
	return t_i
}