package main

import "fmt"

func permutations(nums []int) [][]int {
	var length = len(nums)
	if length == 1{
		return [][]int{nums}
	}
	var result [][]int
	for i:=0;i<length;i++{
		tmp := make([]int, length-1)
		copy(tmp[0:], nums[0:i])
		copy(tmp[i:], nums[i+1:])
		sub := permutations(tmp)
		for _,s := range sub{
			result = append(result, append(s,nums[i]))
		}
	}
	return result
}

func main()  {
	var nums = []int{1,2,3}
	fmt.Println(permutations(nums))
}