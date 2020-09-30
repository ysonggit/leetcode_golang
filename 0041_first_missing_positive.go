/**
 * ideally, A[i] = i+1 swap A[A[i]-1] with A[i]
 * maintain: A[A[i] ]= i by keep on swapping numbers
 *     0 1  2 3
 *    [3 4 -1 1]
 * i=0 A[i]=3 A[A[i]-1] = -1
 *    [-1 4 3 1]
 * i=1 A[i]=4 A[A[i]-1] = 1
 *    [-1 1 3 4]
 * i=2 A[i]=3 A[A[i]-1] = 4
 *    return i+1 such that A[i] != i+1
 */
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] < n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}