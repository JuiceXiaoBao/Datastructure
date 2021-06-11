package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nikename string
	pre      *HeroNode //表示指向前一个结点
	next     *HeroNode //表示指向下一个结点
}

//给双向链表插入一个结点
//编写第一种插入方法,在双向链表的最后加入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//1.先找到该链表的最后这个结点
	//2.创建一个辅助结点[跑龙套，帮忙]
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next //让temp不断的指向下一个结点
	}
	//3.将newHeroNode加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

//给链表插入一个结点
//编写第二种插入方法,根据no的编号从小到大插入...
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//1.先找到适当的结点
	//2.创建一个辅助结点[跑龙套，帮忙]
	temp := head
	flag := true
	//让插入结点的no
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			//说明newHeroNode就应该插入到temp后面
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起，已经存在no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}

}

//显示链表的所有节点信息
//仍然使用单向链表的使用方式
//func ListHeroNode(head *HeroNode) {
//	//1.创建一个辅助结点【跑龙套】
//	temp:=head
//
//	//先判断该链表是不是一个空的链表
//	if temp.next==nil {
//		fmt.Println("空空如也...")
//		return
//	}
//
//	//2.遍历这个链表
//	for {
//		fmt.Printf("[%d,%s,%s]==>",temp.next.no,temp.next.name,temp.next.nikename)
//		//判断是否链表后
//		temp = temp.next
//		if temp.next==nil {
//			break
//		}
//	}
//	fmt.Println()
//}

//显示链表的所有节点信息
//仍然使用单向链表的使用方式
func ListHeroNode2(head *HeroNode) {
	//1.创建一个辅助结点【跑龙套】
	temp := head

	//先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也...")
		return
	}

	//2.让temp定位到双向链表的最后结点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	//2.遍历这个链表
	for {
		fmt.Printf("[%d,%s,%s]==>", temp.no, temp.name, temp.nikename)
		//判断是否到了链表的头部
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
	fmt.Println()
}

//删除一个结点[双向链表删除结点]
func DelHerNode(head *HeroNode, id int) {
	temp := head
	flag := false
	//找到要删除结点的no，和temp的下一个结点的no比较
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next //ok
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("sorry,要删除的id不存在！")
	}
}
func main() {

	//先创建一个头结点,采用默认值
	head := &HeroNode{}

	//2.创建一个新的HeroNode
	hero := &HeroNode{
		no:       1,
		name:     "宋江",
		nikename: "及时雨",
		next:     nil,
	}

	hero1 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nikename: "玉麒麟",
		next:     nil,
	}

	hero2 := &HeroNode{
		no:       3,
		name:     "林冲",
		nikename: "豹子头",
		next:     nil,
	}

	InsertHeroNode(head, hero)
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	ListHeroNode2(head)
}
