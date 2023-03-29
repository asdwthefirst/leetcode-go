/**
  @author: jiangxi
  @since: 2022/11/11
  @desc: //TODO
**/
package stackqueue

type MyQueue struct {
	inStack, outStack []int
}

func Constructor232() MyQueue {
	return MyQueue{}

}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)

}

func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		for i := len(this.inStack) - 1; i >= 0; i-- {
			this.outStack = append(this.outStack, this.inStack[i])
			this.inStack = this.inStack[:i]
		}
	}
	res := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return res
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		for i := len(this.inStack) - 1; i >= 0; i-- {
			this.outStack = append(this.outStack, this.inStack[i])
			this.inStack = this.inStack[:i]
		}
	}
	res := this.outStack[len(this.outStack)-1]
	return res
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
