package main

func removeDuplicates(nums []int) int {
	var n int
	for i:=0;i<len(nums)-1;i++{
		if nums[i] != nums[i+1]{
			nums[n] = nums[i+1]
			n++
		}
	}

	return n
}
