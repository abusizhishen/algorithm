package main

func maxProduct(words []string) int {
	var max int
	for i:=0;i<len(words)-1;i++{
		for j:=i+1;j<len(words);j++{
			if hasSameWord(words[i],words[j]){
				continue
			}

			num := len(words[i])*len(words[j])
			if num > max{
				max = num
			}
		}
	}

	return max
}

//列出所有的两两组合
//过滤含有相同字母的组合

// 求每个组合length的 乘机
//求最大乘机


//位运算、map
func hasSameWord(a,b string)bool{

	for i:=0;i<len(a);i++{
		for j:=0;j<len(b);j++{
			if a[i] == b[j]{
				return true
			}
		}
	}

	return false
}

