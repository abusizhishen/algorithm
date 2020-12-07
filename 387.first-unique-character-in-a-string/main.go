package main

func main(){
	var s = "loveleetcode"
	println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	var arr [26]int
	for i,k := range s{
		arr[k-'a'] = i
	}

	for i,k := range s{
		if arr[k-'a'] == i{
			return i
		}else {
			arr[k-'a'] = -1
		}
	}

	return -1
}