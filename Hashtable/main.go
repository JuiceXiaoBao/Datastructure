package main

import (
	"fmt"
	"os"
)

//定义emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func showMe(emp *Emp) {
	fmt.Printf("链表%d 找到该雇员%d\n", emp.Id%7, emp.Id)
}

//定义EmpLink
//这里我们的EmpLink不带表头，即第一个结点就存放雇员
type EmpLink struct {
	Head *Emp
}

//1.添加员工的方法，保证添加时，编号从小到大
func (em *EmpLink) Insert(emp *Emp, linkNo int) {
	cur := em.Head     //辅助指针
	var pre *Emp = nil //这是一个辅助指针pre在cur前面
	//如果当前的Emplink就是一个空链表
	if cur == nil {
		em.Head = emp
		return
	}
	//如果不是一个空链表，给emp找到对应的位置并插入
	//让cur和emp比较，然后让pre保持在cur前面
	for {
		if cur != nil {
			//比较
			if cur.Id > emp.Id {
				//找到位置
				break
			} else if cur.Id == emp.Id {
				fmt.Printf("你输入的Id在%02d链表中已经存在，请重新设置！\n", linkNo)
				return
			}
			pre = cur //保证同步
			cur = cur.Next
		} else {
			break
		}
	}

	//退出时，我们看下是否将emp添加到链表最后,此步不懂可回看视频！
	pre.Next = emp
	emp.Next = cur

}

//显示链表的信息
func (em *EmpLink) ShowLink(nu int) {
	if em.Head == nil {
		fmt.Printf("链表%d 为空\n", nu)
		return
	}

	//遍历当前的链表，并显示数据
	cur := em.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", nu, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}

	}
	fmt.Println() //换行处理
}

//根据id查找指定的雇员
func (em *EmpLink) FindByte(id int) *Emp {
	cur := em.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

//根据id删除指定的雇员
func (em *EmpLink) Deletehaxi(id int) {
	cur := em.Head
	whx := em.Head.Next //删除头结点时还需要多一个whx这个辅助变量来帮忙
	var pre *Emp
	for {
		if cur != nil && cur.Id == id {
			if id == em.Head.Id {
				em.Head = whx //删除头结点时还需要多一个whx这个辅助变量来帮忙
				fmt.Printf("删除头结点id为%02d成功\n", id)
				return
			} else {
				pre.Next = cur.Next
				fmt.Printf("删除id为%02d成功\n", id)
				return
			}

		} else if cur == nil {
			fmt.Printf("id为%d的编号没有找到\n", id)
			return
		}
		pre = cur
		cur = cur.Next
	}

}

//定义hashtable，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

//给HashTable编写Insert雇员的方法
func (ha *HashTable) Insert(emp *Emp) {
	//使用散列函数,确定将该雇员添加到哪个链表
	linkNo := ha.HashFun(emp.Id)
	ha.LinkArr[linkNo].Insert(emp, linkNo)
}

//编写方法，显示哈希表所有雇员
func (ha *HashTable) showAll() {
	for i := 0; i < len(ha.LinkArr); i++ {
		ha.LinkArr[i].ShowLink(i)
	}
}

//编写一个散列方法
func (ha *HashTable) HashFun(id int) int {
	return id % 7
}

//编写一个方法完成查找
func (ha *HashTable) FindById(id int) *Emp {
	//使用散列函数,确定该雇员在哪个链表
	linkNo := ha.HashFun(id)
	return ha.LinkArr[linkNo].FindByte(id)
}

func (ha *HashTable) Deletehash(id int) {
	//使用散列函数,确定该雇员在哪个链表
	linkNo := ha.HashFun(id)
	ha.LinkArr[linkNo].Deletehaxi(id)
}

func main() {

	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("==============雇员系统菜单==============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show  表示显示雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("delete 表示删除雇员")
		fmt.Println("exit 表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
				Next: nil,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.showAll()
		case "find":
			fmt.Println("请输入id号")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d的雇员不存在！\n", id)
			} else {
				showMe(emp)
			}
		case "delete":
			fmt.Println("请输入id号")
			fmt.Scanln(&id)
			hashtable.Deletehash(id)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误！")

		}
	}
}
