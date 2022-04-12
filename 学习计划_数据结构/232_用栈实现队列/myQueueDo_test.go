package myQueueDo

import "testing"

// 单数组行为
//type MyQueue struct {
//	Stack []int
//}
//
//func Constructor() MyQueue {
//	return MyQueue{make([]int, 0)}
//
//}
//
//func (this *MyQueue) Push(x int) {
//	this.Stack = append(this.Stack, x)
//}
//
//func (this *MyQueue) Pop() int {
//	num := this.Peek()
//	if len(this.Stack) > 1 {
//		this.Stack = this.Stack[1:]
//	} else {
//		this.Stack = make([]int, 0)
//	}
//	return num
//}
//
//func (this *MyQueue) Peek() int {
//	if len(this.Stack) < 1 {
//		return 0
//	}
//	num := this.Stack[0]
//	return num
//}
//
//func (this *MyQueue) Empty() bool {
//	return len(this.Stack) < 1
//}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
func Test232(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	t.Log(obj)
	t.Log(obj.Pop())
	t.Log(obj)
	t.Log(obj.Peek())
	t.Log(obj)
	t.Log(obj.Empty())
}

// 双栈组合
type MyQueue struct {
	inStack, outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

// 数据入栈
func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

// 将入栈的数据反向塞给出栈
func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		// 清空结尾数据（性能提升）
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

func (q *MyQueue) Pop() int {
	// 如果出栈没值，将入栈的数据反向塞给出栈
	if len(q.outStack) == 0 {
		q.in2out()
	}
	// 取出栈最后一个值
	x := q.outStack[len(q.outStack)-1]
	// 删除出栈最后一个值（性能提升）
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}

func (q *MyQueue) Peek() int {
	// 如果出栈没值，将入栈的数据反向塞给出栈
	if len(q.outStack) == 0 {
		q.in2out()
	}
	// 返回出栈最后一个值
	return q.outStack[len(q.outStack)-1]
}

// 出入栈都没值
func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}
