func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	sum := 0
	for i := 0; i < n; i++ {
		for l := 1; i+l <= n; l += 2 {
			sum += sumArray(arr[i : i+l])
		}
	}
	return sum
}

func sumArray(arr []int) int {
	total := 0
	for _, i := range arr {
		total += i
	}
	return total
}