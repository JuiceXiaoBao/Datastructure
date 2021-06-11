package main

import "fmt"

type Boy struct {
	No   int  //编号
	Next *Boy //指向下一个小孩的指针
}

//编写一个函数，构成单向的环形链表
//num:表示小孩的个数
//*Boy：返回该环形的链表的第一个小孩的指针
func Addboy(num int) *Boy {
	first := &Boy{}  //空结点
	curBoy := &Boy{} //空结点
	//判断
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}
	//循环的构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No:   i,
			Next: nil,
		}
		//分析构成循环链表，需要一个辅助指针来帮忙
		//1.因为第一个小孩比较特殊
		if i == 1 { //第一个小孩
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first //构成环形链表
		}
	}
	return first
}

//显示单向的环形链表[遍历]
func ShowBoy(first *Boy) {
	//环形链表为空时
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩...")
		return
	}
	//创建一个指针，帮助遍历
	curBoy := first
	for {
		fmt.Printf("小孩编号为=%d ->\n", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}

}

/*
设编号为1，2，_ n 的n个人围坐一圈，约定编号为k（1<=k<=n）的人从1开始报数
数到m的那个人出列，它的下一代又从1开始报数，数到m的那个人又出列，依次类推
直到所有人出列为止，由此产生一个出队编号的序列
*/

func PlayGame(first *Boy, startNo int, countNum int) {
	//1.空的链表我们单独处理
	if first.Next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}

	//
	//2.需要定义辅助指针，帮助我们删除小孩
	tail := first
	//3.tail执行环形链表的最后一个小孩，这个非常重要
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	//4.让first移动到startNo【后面我们删除小孩，就以first为准】
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	//5.开始数countNum，然后就删除first指向的小孩
	for {
		//开始数countNum-1次
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d出圈 ->", first.No)
		//删除first执行的小孩
		first = first.Next
		tail.Next = first
		//判断如果tail==first，说明圈中只有一个小孩
		if tail == first {
			break
		}
	}
	fmt.Printf("最后一个小孩编号为%d 出圈->", first.No)

}

func main() {
	boy := Addboy(5)
	ShowBoy(boy)
	PlayGame(boy, 2, 3)
}
