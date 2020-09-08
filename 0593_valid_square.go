import "sort"

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	coords := [][]int{p1, p2, p3, p4}
	sides_lens := make(map[int]int, 0)
	sides := []int{}
	for i := 0; i < len(coords); i++ {
		a := coords[i]
		for j := i + 1; j < len(coords); j++ {
			b := coords[j]
			side_len := dist(a, b)
			sides = append(sides, side_len)
			if _, ok := sides_lens[side_len]; ok {
				sides_lens[side_len] += 1
			} else {
				sides_lens[side_len] = 1
			}
		}
	}
	sort.Ints(sides)
	return sides_lens[sides[0]] == 4 && sides_lens[sides[len(sides)-1]] == 2
}

func dist(p1 []int, p2 []int) int {
	// given assumptions, no overflow handling needed
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}