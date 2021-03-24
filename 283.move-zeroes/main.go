package main

func moveZeroes(nums []int)  {
	var j int
	for i:=0;i<len(nums);i++{
		if nums[i] == 0 {
			continue
		}

		nums[i],nums[j] = nums[j],nums[i]
		j++

	}
}