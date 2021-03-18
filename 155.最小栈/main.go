package main

//两个栈
//数据栈用来pop push
//最小栈用来保存数据栈对应num对应的最小值

type MinStack struct {
	stack []int
	minStack []int
	min int
	hasMin bool
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack, val)

	if !this.hasMin {
		this.minStack = append(this.minStack, val)
		this.hasMin = true
		this.min = val
		return
	}

	if val < this.min {
		this.min = val
	}

	this.minStack = append(this.minStack, this.min)
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
	if len(this.minStack) != 0 {
		this.min = this.minStack[len(this.minStack)-1]
		return
	}
	this.hasMin = false
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
