package main

import (
	"fmt"
)


//如果之前的和小于零，扔掉之前的和
func main()  {
	var nums = []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	var max = nums[0]
	var sum = max
	for i:=1;i<len(nums);i++{
		if sum < 0{
			sum = 0
		}

		sum+=nums[i]
		if sum > max{
			max = sum
		}
	}

	return max
}
