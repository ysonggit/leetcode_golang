import "container/heap"

type Minheap [][]int

func (m Minheap) Len() int {
	return len(m)
}

func (m Minheap) Less(i, j int) bool {
	return m[i][0]+m[i][1] < m[j][0]+m[j][1]
}

func (m Minheap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *Minheap) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *Minheap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

// brute force method only beats 5% in time and space
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	res, hp := [][]int{}, &Minheap{}
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return res
	}
	heap.Init(hp)
	for _, i := range nums1 {
		heap.Push(hp, []int{i, nums2[0], 0})
	}
	cnt := 0
	for len(*hp) > 0 {
		min := (*hp)[0]
		heap.Pop(hp)
		res = append(res, min[:2])
		cnt++
		if cnt == k {
			break
		}
		if min[2]+1 < len(nums2) {
			heap.Push(hp, []int{min[0], nums2[min[2]+1], min[2] + 1})
		}
	}
	return res
}