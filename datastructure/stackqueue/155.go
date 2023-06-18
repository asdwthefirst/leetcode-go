/*
*

	@author: jiangxi
	@since: 2022/11/11
	@desc: //TODO

*
*/
package stackqueue

type MinStack struct {
	stack    []int
	minStack []int //minstack和stack同数量增长，minstack保存当前最小值
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	min := val
	if len(this.minStack) != 0 {
		min = this.minStack[len(this.minStack)-1]
		if val < min {
			min = val
		}
	}
	this.minStack = append(this.minStack, min)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]

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
