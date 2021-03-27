package main

func singleNumber(nums []int) int {
	var once = map[int]int{}
	for _,num := range nums{
		once[num]++
	}

	for num,count := range once{
		if count == 1 {
			return num
		}
	}

	return 0
}
