package main

func climbStairs(n int) int {
	if n < 3{
		return n
	}
	var step1 = 1
	var step2 = 2
	for i:=3;i<n;i++{
		temp := step2+step1
		step1,step2 = step2,temp
	}

	return step2+step1
}
