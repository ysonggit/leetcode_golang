func restoreString(s string, indices []int) string {
    n := len(s)
    r := make([]rune, n)
    sr := []rune(s)
    for i:=0; i<n; i++ {
        j := indices[i]
        r[j] = sr[i]
    }
    return string(r)
}