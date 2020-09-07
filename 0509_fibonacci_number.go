func fib(N int) int {
	if N < 2 {
		return N
	}
	f_i_2, f_i_1, f_i := 0, 1, 0
	for i := 2; i <= N; i++ {
		f_i = f_i_1 + f_i_2
		f_i_2 = f_i_1
		f_i_1 = f_i
	}
	return f_i
}