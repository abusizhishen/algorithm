package main

import "fmt"

func longestCommonPrefix(ss []string) string {
	if len(ss) <1{
		return ""
	}

	var i,j int
	var sub string
end:
	for {
		for j,sub = range ss{
			if len(sub)==i{
				break end
			}

			if j == 0{
				continue
			}

			if sub[i] != ss[j-1][i]{
				break end
			}
		}
		i++
	}

	if i == 0{
		return ""
	}

	return ss[0][:i]
}

func main() {
	var ss = []string{"abc","ab","abde"}
	fmt.Println(longestCommonPrefix(ss))
}