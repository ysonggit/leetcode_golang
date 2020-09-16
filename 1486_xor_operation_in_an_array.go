func xorOperation(n int, start int) int {
    if n == 0 {
        return 0
    }
    res := start
    for i:=1; i<n; i++ {
        res ^= start + 2 * i
    }
    return res
}