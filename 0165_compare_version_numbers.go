func compareVersion(version1 string, version2 string) int {
	n1, n2 := len(version1), len(version2)
	i, j := 0, 0
	for i < n1 || j < n2 {
		v1, v2 := 0, 0
		for i < n1 && version1[i] != '.' {
			v1 = v1*10 + int(rune(version1[i])-'0')
			i++
		}
		for j < n2 && version2[j] != '.' {
			v2 = v2*10 + int(rune(version2[j])-'0')
			j++
		}
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		} else { // pass .
			i++
			j++
		}
	}
	return 0
}