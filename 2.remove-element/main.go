package main

import "fmt"

func main() {
	var nums = []int{1,2,3,6,7,3,4,5,6}
	var val = 3
	l := removeElement(nums,val)
	fmt.Println(nums[:l])
}

func removeElement(nums []int, val int) int {
	var l int
	for i:=0;i<len(nums);i++{
		if nums[i] != val{
			nums[l] = nums[i]
			l++
		}
	}

	return l
}
