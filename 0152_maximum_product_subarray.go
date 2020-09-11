func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	local_max, local_min, glob_max := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		if cur < 0 {
			local_max, local_min = local_min, local_max
		}
		if cur < local_max*cur {
			local_max *= cur
		} else {
			local_max = cur
		}
		if cur < local_min*cur {
			local_min = cur
		} else {
			local_min *= cur
		}
		if local_max > glob_max {
			glob_max = local_max
		}
	}
	return glob_max
}