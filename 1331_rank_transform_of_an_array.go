import "sort"

func arrayRankTransform(arr []int) []int {
	m := make(map[int]int)
	arr_sort := make([]int, len(arr))
	copy(arr_sort, arr) // must do this!
	sort.Ints(arr_sort)
	res := make([]int, 0)
	rank := 1
	for _, num := range arr_sort {
		if _, ok := m[num]; ok {
			continue
		}
		m[num] = rank
		rank++
	}
	for _, num := range arr {
		res = append(res, m[num])
	}
	return res
}