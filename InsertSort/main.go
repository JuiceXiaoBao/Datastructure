package main

import "fmt"

func InsertSort(arr *[5]int) {

	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] //数据后移
			insertIndex--
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次插入后数组为：%d\n", i, *arr)

	}
}

//完成第一次,给第二个元素找到合适的位置并插入
//insertVal:=arr[1]
//insertIndex:=1-1
//
////从大到小
//for insertIndex>=0 && arr[insertIndex]<insertVal {
//	arr[insertIndex+1]= arr[insertIndex]  //数据后移
//	insertIndex--
//}
////插入
//if insertIndex+1!=1 {
//	arr[insertIndex+1] = insertVal
//}
//fmt.Println("第一次插入后",*arr)

//}

func main() {
	arr := [5]int{23, 0, 12, 56, 34}
	InsertSort(&arr)
}
