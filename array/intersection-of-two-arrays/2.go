package main

import "fmt"

//两个有序数组求交集
func intersection2(a,b []int)[]int  {
	var i,j = len(a),len(b)
	var m,n,k int
	for m<i&&n<j{
		if a[m]<b[n]{
			m++
			continue
		}else if a[m]>b[n]{
			n++
			continue
		}else {
			a[k] = a[m]
			k++
			m++
			n++
		}
	}

	return a[:k]
}

func main() {
	var a = []int{1,2,5,7,9}
	var b = []int{0,2,3,4,9}

	fmt.Println(intersection2(a,b))
}
