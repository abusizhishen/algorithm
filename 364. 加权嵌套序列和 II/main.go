package main


//step 1 层级遍历，记录当前层次和最深层次
//step 2 循环计算每层权重和
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 
 */

type NestedInteger struct {}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {return false}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int { return 0}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger { return nil}


func depthSumInverse(nestedList []*NestedInteger) int {
	trees,deep := getTrees(nestedList)
	if trees == nil {
		return 0
	}

	var sum int
	for _,tree := range trees{
		sum += tree.Int*(deep+1-tree.Level)
	}

	return sum
}

type Tree struct{
	Int int
	Level int
}

func getTrees(nestedList []*NestedInteger) ([]Tree,int){
	if nestedList == nil || len(nestedList) == 0 {
		return nil,0
	}

	var trees []Tree
	var deep int

	for nestedList != nil {
		deep++
		var temp []*NestedInteger
		var tree Tree
		tree.Level = deep

		for _,node := range nestedList {
			if node == nil {
				continue
			}

			tree.Int += node.GetInteger()
			list := node.GetList()
			if list == nil {
				continue
			}

			temp = append(temp, list...)
		}

		if tree.Int != 0 {
			trees = append(trees, tree)
		}
		nestedList = temp
	}

	return trees,deep
}
