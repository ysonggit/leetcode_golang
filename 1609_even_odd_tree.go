	func isEvenOddTree(root *TreeNode) bool {
		if root == nil {
			return false
		}
		queue := []*TreeNode{root}
		level_idx := 0
		for len(queue) > 0 {
			level_vals := []int{}
			level_size := len(queue)
			for i := 0; i < level_size; i++ {
				cur := queue[i]
				if level_idx%2 == 0 {
					if cur.Val%2 != 1 {
						return false
					}
				} else {
					if cur.Val%2 != 0 {
						return false
					}
				}
				level_vals = append(level_vals, cur.Val)
				if cur.Left != nil {
					queue = append(queue, cur.Left)
				}
				if cur.Right != nil {
					queue = append(queue, cur.Right)
				}
			}
			queue = queue[level_size:]
			if level_idx%2 == 0 {
				if !increasingList(level_vals) {
					return false
				}
			} else {
				if !decreasingList(level_vals) {
					return false
				}
			}
			level_idx++
		}
		return true
	}

	func increasingList(list []int) bool {
		for i := 0; i < len(list)-1; i++ {
			if list[i] >= list[i+1] {
				return false
			}
		}
		return true
	}

	func decreasingList(list []int) bool {
		for i := 0; i < len(list)-1; i++ {
			if list[i] <= list[i+1] {
				return false
			}
		}
		return true
	}