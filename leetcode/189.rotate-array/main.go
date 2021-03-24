package main

import "fmt"

func rotate(nums []int, k int)  {
	k %= len(nums)

	if k == 0{
		return
	}
	var temp = make([]int,k)
	for i:=0;i<k;i++{
		temp[i] = nums[i]
	}

	start := len(nums)-1
	end := k
	for start>=end {
		index := (start+k)%len(nums)
		nums[index] = nums[start]
		start--
	}

	for i:=k-1;i>=0;i--{
		index := (start+k)%len(nums)
		nums[index] = temp[start]
		start--
	}
}

func rotate2(nums []int, k int)  {
	if len(nums)<=1{
		return
	}
	k %=len(nums)
	var temp int
	for i:=0;i<k;i++{
		temp = nums[len(nums)-1]
		for j:=len(nums)-1;j>0;j--{
			nums[j] = nums[j-1]
		}
		nums[0] = temp
	}
}

func main() {
	var nums = []int{1,2,3,4,5}
	rotate2(nums,2)
	fmt.Println(nums)
}