package main

import "fmt"

func main() {
	var nums =[]int{1,2,3,4,5,6,7}
	var k = 3
	var l = len(nums)-1
	k %= len(nums)

	reverse(nums,0,l)
	reverse(nums,0,k-1)
	reverse(nums,k,l)
	fmt.Println(nums)
}

func reverse(nums []int, start,end int)  {
	for start<=end{
		nums[start],nums[end] = nums[end],nums[start]
		start++
		end--
	}
}
