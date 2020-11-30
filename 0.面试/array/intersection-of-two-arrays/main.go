package main

import "fmt"

func main() {
	var a = []int{1,2,3,6,7,8}
	var b = []int{1,4,5,6,3}


	fmt.Println(intersection(a,b))
}

func intersection(a,b []int)[]int  {
	var m = map[int]int{}
	for _,num := range a{
		m[num]+=1
	}

	var k int
	for _,num := range b{
		if m[num]>0{
			m[num]--
			a[k]=num
			k++
		}
	}

	return a[0:k]
}
