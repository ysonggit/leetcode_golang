func maxPower(s string) int {
	max_len := 1
	cur_len := 1
	for fast := 1; fast < len(s); fast++ {
		if s[fast] != s[fast-1] {
			if max_len < cur_len {
				max_len = cur_len
			}
			cur_len = 1
		} else {
			cur_len++
		}
	}
	if max_len < cur_len { // ccc
		return cur_len
	}
	return max_len
}