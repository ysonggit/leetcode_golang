func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	i, j, k := 0, 0, 0
	res := []int{}
	for i < len(arr1) && j < len(arr2) && k < len(arr3) {
		if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
			res = append(res, arr1[i])
			i, j, k = i+1, j+1, k+1
		} else {
			if arr1[i] < arr2[j] {
				i++
			} else if arr2[j] < arr3[k] {
				j++
			} else if arr3[k] < arr1[i] {
				k++
			}
		}
	}
	return res
}