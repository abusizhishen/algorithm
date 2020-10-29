package main


//伴鱼 矩阵行列转换
import "fmt"

func main() {
	var data = [][]int{
		[]int{1,2,3},
		[]int{1,2,3},
		[]int{1,2,3},
	}
	print(data)
	print(reverse(data))
}

func reverse(arrs [][]int) (result [][]int) {
	var height = len(arrs)
	if height==0{
		return
	}

	var width = len(arrs[0])
	if width == 0{
		return
	}
	result = make([][]int,width)
	for i:=0;i<width;i++{
		result[i] = make([]int,height)
	}

	for i:=0;i<height;i++{
		for j:=0;j<width;j++{
			result[j][i] = arrs[i][j]
		}
	}

	return result
}

func print(arrs [][]int)  {
	for i:=0;i<len(arrs);i++{
		for j:=0;j<len(arrs[i]);j++{
			fmt.Printf("%d,", arrs[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
