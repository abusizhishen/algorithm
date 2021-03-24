package main

//向中间夹逼
func maxArea(height []int) int {
	var maxArea int
	var i,j =0,len(height)-1
	for i<j{
		x := j-i
		y := height[i]
		if height[i]>height[j]{
			y = height[j]
			j--
		}else{
			i++
		}

		maxArea = max(maxArea,x*y)
	}

	return maxArea
}

func max(a,b int)int{
	if a>b {
		return a
	}

	return b
}
