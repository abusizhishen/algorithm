package main

import "fmt"

func main() {
	nums := []int{1,2,3}
	fmt.Println(permute(nums))
}

//递归求解
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return [][]int{nums}
	}

	var result [][]int
	for i:=0;i<len(nums);i++{
		var tmp = []int{nums[i]}
		var subNums []int
		subNums = append(subNums,nums[0:i]...)
		subNums = append(subNums,nums[i+1:]...)

		for _, sub := range permute(subNums){
			result = append(result, append(tmp,sub...))
		}
	}

	return result
}