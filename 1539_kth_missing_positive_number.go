func findKthPositive(arr []int, k int) int {
    expect := 1
    n := len(arr)
    for i:= 0; i < n; i++ {
        for {
            if expect == arr[i] {
                break
            }else{
                k--
                if k==0 {
                    return expect
                }
                expect++
            }
        }
        expect++
    }
    return arr[n-1]+k
}