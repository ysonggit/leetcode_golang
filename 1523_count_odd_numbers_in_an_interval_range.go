func countOdds(low int, high int) int {
    dist := high - low
    odds := dist / 2
    // if high and low are both even, that's all odds 
    // if there 1 or 2 odds between high and low, need to add 1 more
    if low & 1 == 1 || high & 1== 1 {
        odds++
    }
    return odds
}