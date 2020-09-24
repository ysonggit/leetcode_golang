func findTheDifference(s string, t string) byte {
	alph := [26]int{}
	for _, c := range s {
		alph[c-97]++
	}
	var res byte
	for _, c := range t {
		if alph[c-97] == 0 {
			res = byte(c)
			break
		} else {
			alph[c-97]--
		}
	}
	for _, i := range alph {
		if i < 0 {
			res = byte(i + 97)
			break
		}
	}
	return res
}