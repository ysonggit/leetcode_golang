import "strconv"

func largestTimeFromDigits(arr []int) string {
	all_perms := permutations(arr)
	max_time := ""
	for _, t := range all_perms {
		hh, _ := strconv.Atoi(t[0:2])
		mm, _ := strconv.Atoi(t[2:])
		if hh < 24 && mm < 60 {
			if len(max_time) == 0 {
				max_time = t[0:2] + ":" + t[2:]
			} else {
				max_time_hh, _ := strconv.Atoi(max_time[0:2])
				max_time_mm, _ := strconv.Atoi(max_time[3:])
				if hh*60+mm > max_time_hh*60+max_time_mm {
					max_time = t[0:2] + ":" + t[2:]
				}
			}
		}
	}
	return max_time
}

func permutations(arr []int) []string {
	s := ""
	for _, i := range arr {
		s += strconv.Itoa(i)
	}
	all_perms := make([]string, 0)
	perms(s, "", &all_perms)
	return all_perms
}

func perms(s string, cur string, arr *[]string) {
	if len(s) == 0 {
		*arr = append(*arr, cur)
	}
	for i := 0; i < len(s); i++ {
		rest := s[0:i] + s[i+1:]
		perms(rest, cur+string(s[i]), arr)
	}
}