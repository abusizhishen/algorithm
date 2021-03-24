package main

func plusOne(digits []int) []int {
	var pre = 1
	for i:=len(digits)-1;i>=0;i--{
		digits[i]+=pre
		if digits[i] > 9{
			digits[i] -= 10
		}else{
			pre=0
			break
		}
	}

	if pre == 1{
		var nums = make([]int,0, len(digits)+1)
		nums = append(nums,1)
		nums = append(nums,digits...)
		return nums
	}

	return digits
}
