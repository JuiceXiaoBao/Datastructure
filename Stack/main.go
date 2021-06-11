package main

import (
	"errors"
	"fmt"
)

//使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop int //表示栈最大可以存放的个数
	Top    int //表示栈顶，因为栈顶固定，因此我们直接使用Top
	arr    [5]int
}

//入栈
func (st *Stack) Push(val int) (err error) {

	//先判断栈是否满了
	if st.Top == st.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	st.Top++
	st.arr[st.Top] = val
	return
}

//出栈
func (st *Stack) Pop() (val int, err error) {
	//先判断栈是否为空
	if st.Top == -1 {
		return 0, errors.New("stack empty")
	}

	//先取值，再st.Top--
	val = st.arr[st.Top]
	st.Top--
	return val, nil
}

//遍历栈,注意需要从栈顶开始遍历
func (st *Stack) List() {
	//先判断栈是否为空
	if st.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	for i := st.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\t", i, st.arr[i])
	}
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
		arr:    [5]int{},
	}
	//入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	//显示
	stack.List()
	val, err := stack.Pop()
	val, err = stack.Pop()
	val, err = stack.Pop()
	val, err = stack.Pop()
	val, err = stack.Pop()
	val, err = stack.Pop()
	fmt.Println()
	fmt.Printf("被弹出的值为%d,错误为%v\n", val, err)
	stack.List()

}
