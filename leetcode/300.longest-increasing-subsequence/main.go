package main

import "fmt"

func lis(nums []int) int {
	if len(nums)==0{
		return 0
	}
	var dp = make([]int,len(nums))
	var result int = 1
	for i:=0;i<len(nums);i++{
		dp[i] = 1
		for j:=0;j<i;j++{
			if nums[j]<nums[i]{
				dp[i] = max(dp[j]+1,dp[i])
			}
		}

		result = max(result, dp[i])
	}

	return result
}

func max(a,b int) int {
	if a>=b{
		return a
	}

	return b
}

func main() {
	var nums = []int{1,2,7,3,0,1,2,3}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	return result
}