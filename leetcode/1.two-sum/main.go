package main

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))
	var i=0
	for  {
		if idx,ok := m[target-nums[i]];ok{
			return []int{idx,i}
		}else{
			m[nums[i]] = i
			i++
		}
	}
}
