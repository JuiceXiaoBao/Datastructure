package main

import (
	"fmt"
	"math/rand"
	"time"
)

//编写函数selectSort完成选择排序
func SelectSort(arr *[80000]int) {
	//1.假设arr【0】最大值
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxindex := j
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxindex = i
			}
		}
		//交换
		if maxindex != j {
			arr[j], arr[maxindex] = arr[maxindex], arr[j]
		}
		//fmt.Printf("第%d次 %v\n", j+1, *arr)
	}
}

/*
		max=arr[1]
		maxindex=1
		for i:=2;i<len(arr);i++ {
			if max<arr[i] {
				max = arr[i]
				maxindex=i
			}
		}
		//交换
		if maxindex!=1 {
			arr[1],arr[maxindex]=arr[maxindex],arr[1]
		}
		fmt.Println("第二次",*arr)

		max=arr[2]
		maxindex=2
		for i:=3;i<len(arr);i++ {
			if max<arr[i] {
				max = arr[i]
				maxindex=i
			}
		}
		//交换
		if maxindex!=2 {
			arr[2],arr[maxindex]=arr[maxindex],arr[2]
		}
		fmt.Println("第三次",*arr)

		max=arr[3]
		maxindex=3
		for i:=4;i<len(arr);i++ {
			if max<arr[i] {
				max = arr[i]
				maxindex=i
			}
		}
		//交换
		if maxindex!=3 {
			arr[3],arr[maxindex]=arr[maxindex],arr[3]
		}
		fmt.Println("第四次",*arr)
	}

*/
func main() {
	//定义一个数组
	//arr := [5]int{10, 34, 19, 100, 80}
	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}
	start := time.Now().Unix()
	SelectSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("选择排序耗时=%d秒", end-start)
}
