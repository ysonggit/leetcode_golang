func strWithout3a3b(A int, B int) string {
	res := ""
	return recursiveAppend(A, B, res)
}

func recursiveAppend(A int, B int, cur string) string {
	if A == 0 || B == 0 {
		for A > 0 {
			cur += "a"
			A--
		}
		for B > 0 {
			cur += "b"
			B--
		}
	} else if A == B {
		cur += "ab"
		cur = recursiveAppend(A-1, B-1, cur)
	} else if A > B {
		cur += "aab"
		cur = recursiveAppend(A-2, B-1, cur)
	} else if A < B {
		cur += "bba"
		cur = recursiveAppend(A-1, B-2, cur)
	}
	return cur
}