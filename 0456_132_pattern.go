/*
https://leetcode.com/problems/132-pattern/discuss/902824/C%2B%2B-Stack-O(N)-Solution-Explained
Given array nums, find nums[i] < nums[k] < nums[j] such that i < j < k

[3,1,4,2]

For tracking nums[i] we can use minValue array where ith value stores 
min value in range 0 to i-1 index

[3,1,1,1]

Now we can loop from end to start and store the next greater number 
of minValue[x] in the stack and check if stack.top < nums[x] then return true
    
minValue[x] (nums[i]) < stack.top (nums[k]) < nums[x] (nums[j])
stack   min_nums
[2 ]    1
*/
func find132pattern(nums []int) bool {
    n := len(nums)
    if n < 3 {
        return false
    }
    min_nums := make([]int, n)
    min_nums[0] = nums[0]
    for i:=1; i<n; i++ {
        if nums[i] < min_nums[i-1] {
            min_nums[i] = nums[i]
        }else{
            min_nums[i] = min_nums[i-1]
        }
    }
    stack := make([]int, 0)
    for i:=n-1; i>=0; i-- {
        
        for len(stack) > 0 && stack[len(stack)-1] <= min_nums[i] {
            stack = stack[:len(stack)-1]
        }
        if len(stack) > 0 && stack[len(stack)-1] < nums[i] {
            return true
        }
        stack = append(stack, nums[i])
        /*
        fmt.Printf("Stack: [")
        for _ , x := range stack {
            fmt.Printf("%d, ", x)
        }
        fmt.Println("]")
        */
    }
    return false  
}
