func majorityElement(nums []int) []int {
	majs, n := make([]int, 0), len(nums)
	if n == 0 {
		return majs
	}
	maj1, maj2 := nums[0], nums[0]
	cnt1, cnt2 := 0, 0
	for _, i := range nums {
		if i == maj1 {
			cnt1++
			continue
		}
		if i == maj2 {
			cnt2++
			continue
		}
		if cnt1 == 0 {
			maj1 = i
			cnt1 = 1
			continue
		}
		if cnt2 == 0 {
			maj2 = i
			cnt2 = 1
			continue
		}
		cnt1--
		cnt2--
	}
	cnt1, cnt2 = 0, 0
	for _, i := range nums {
		if i == maj1 {
			cnt1++
			continue
		}
		if i == maj2 {
			cnt2++
		}
	}
	if cnt1 > n/3 {
		majs = append(majs, maj1)
	}
	if cnt2 > n/3 {
		majs = append(majs, maj2)
	}
	return majs
}