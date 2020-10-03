func findPairs(nums []int, k int) int {
    counter := make(map[int]int)
    for _, n := range nums {
        counter[n]+=1
    }
    res := 0
    for n, freq := range counter {
        if k == 0 {
            if freq > 1 {
               res++ 
            }
        } else { 
            if _, ok := counter[n+k]; ok {
                res++
            }
        }
    }
    return res
}
