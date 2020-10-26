class Solution:
    def checkArithmeticSubarrays(self, nums: List[int], l: List[int], r: List[int]) -> List[bool]:  
        def isArithmetic(arr):
            if len(arr) < 2:
                return False
            if len(arr) == 2:
                return True
            arr.sort()
            diff = arr[1] - arr[0]
            for i in range(2, len(arr)):
                if arr[i] - arr[i-1] != diff:
                    return False
            return True
        
        n = len(r)
        res = [False] * n
        for i in range(n):
            start, end = l[i], r[i]
            sub_nums = nums[start:end+1]
            res[i] = isArithmetic(sub_nums)
        return res
