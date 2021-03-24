package main

import "fmt"

//数组写入index需注意
func isValid(s string) bool {
	var m = map[string]string{
		"]":"[",
		")":"(",
		"}":"{",
	}
	var length = len(s)
	var arr = make([]string,length)
	var index int

	for i:=0;i<length;i++{
		ss := string(s[i])
		if _,ok := m[ss];!ok{
			arr[index] = ss
			index++
			continue
		}

		if i == 0 {
			return false
		}

		if arr[index-1] == m[ss]{
			if index == 0 {
				return false
			}
			index--
		}else{
			return false
		}
	}

	return index==0
}

func match(s1,s2 string)bool{
	return false
}

func main()  {
	var s = "(){}}{"
	fmt.Println(isValid(s))
}
