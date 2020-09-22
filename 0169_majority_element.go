func majorityElement(nums []int) int {
	n := len(nums)
	maj := nums[0]
	if n == 1 {
		return maj
	}
	cnt := 0
	for _, i := range nums {
		if i == maj {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			maj = i
			cnt = 1
		}
	}
	return maj
}