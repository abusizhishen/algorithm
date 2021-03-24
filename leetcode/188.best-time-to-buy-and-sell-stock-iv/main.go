package main

import (
	"fmt"
)

func main() {
	money := maxProfit(2, []int{1,2,4,2,5,7,2,4,9,0})
	fmt.Println(money)
}

func maxProfit(k int, prices []int) int {
	if len(prices) <= 1{
		return 0
	}
	var arr []int
	for i:=1;i<len(prices);i++{
		arr = append(arr, prices[i]-prices[i-1])
	}

	var newArr []int
	var incr bool
	var index int
	for i:=1;i<len(arr);i++{
		if i == 1{
			if arr[1] > arr[0]{
				incr = true
				index = 0
			}
			continue
		}

		
	}

	fmt.Println(arr)
	return 0
}
