package main

func maxProfit(prices []int) int {
	//终止条件
	if len(prices) <=1{
		return 0
	}

	//记录每次的买点
	var max int
	var buy bool
	var i,start int

	for i<len(prices){
		if i==0{
			i++
			continue
		}

		//注意降价了
		if prices[i-1]>prices[i] {
			//买了就赶紧卖出
			if buy{
				max += prices[i-1]-prices[start]
				buy = false
			}
			//升值了
		}else if prices[i-1]<prices[i]{
			//没买就赶紧买
			if !buy{
				buy = true
				start = i-1
			}
		}
		i++
	}

	//再不卖就没机会了
	if buy{
		max += prices[i-1]-prices[start]
	}
	return max
}
