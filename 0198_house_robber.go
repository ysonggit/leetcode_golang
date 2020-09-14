func rob(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    max_pre, max_cur :=0, 0
    for _, cur := range nums {
        max_pre, max_cur = max_cur, max(cur + max_pre, max_cur)
    }
    return max_cur
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
