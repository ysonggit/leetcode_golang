func rand10() int {
    // https://leetcode.com/problems/implement-rand10-using-rand7/discuss/816927/Python-generalised-solution-for-RandM()-using-RandN()
    accp := 7 * 7 - (7 * 7) % 10
    cur := accp
    for cur >= accp {
        cur = (rand7() - 1) * 7 + rand7() - 1
    }
    return cur % 10 + 1 
}
