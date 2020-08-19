func longestPalindrome(s string) string {
    n := len(s)
    if n < 2 {
        return s
    }
    /*
        DP O(n^2)
        P[i][j] : substring s[i]...s[j] is a palindrome
        P[i][j] is true if P[i+1][j-1] is true && s[i]==s[j]
    */
    P := make([][]bool, n)
    for i := range P {
        P[i] = make([]bool, n)
    }
    for i:=0; i<n; i++ {
        for j :=0; j<n; j++ {
            if P[i][j] = true; i==j {
                P[i][j] = false
            }
        }
    }
    maxlen := 1
    start_idx := 0
    for j:=0; j<n; j++ {
        for i:=j; i>=0; i-- {
            P[i][j] = s[i] == s[j] && ((j-i<=1) || P[i+1][j-1])
            if P[i][j] {
                if maxlen < j-i+1 {
                    maxlen = j-i+1
                    start_idx = i
                    //fmt.Printf("start idx: %v, maxlen: %v", start_idx, maxlen)
                }
            }
        }
    }
    return s[start_idx:start_idx+maxlen]
}