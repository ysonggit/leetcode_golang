import "sort"

func triangleNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	cnt := 0
	for i := n - 1; i > 1; i-- {
		l, r := 0, i-1
		for l < r {
			if nums[l]+nums[r] > nums[i] {
				cnt += r - l
				r--
			} else {
				l++
			}
		}
	}
	return cnt
}