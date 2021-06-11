package main

import "fmt"

//自己敲一边快速排序算法
func QuickSort(arr []int, left, right int) {
	if left < right {
		temp := arr[left]
		j := right
		i := left
		for {
			for arr[j] >= temp && i < j {
				j--
			}
			for arr[i] <= temp && i < j {
				i++
			}
			if i >= j {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		arr[left] = arr[i]
		arr[i] = temp
		QuickSort(arr, left, i-1)
		QuickSort(arr, i+1, right)
	}

}

func main() {
	arr := []int{
		1, 8, 1, 2, 9, 5,
	}
	QuickSort(arr, 0, 5)
	fmt.Println(arr)
}

//func quickSort(values []int, left int, right int) {
//
//	if left < right {
//
//		// 设置分水岭
//
//		temp := values[left]
//
//		// 设置哨兵
//
//		i, j := left, right
//
//		for {
//
//			// 从右向左找，找到第一个比分水岭小的数
//
//			for values[j] >= temp && i < j {
//
//				j--
//
//			}
//
//			// 从左向右找，找到第一个比分水岭大的数
//
//			for values[i] <= temp && i < j {
//
//				i++
//
//			}
//
//			// 如果哨兵相遇，则退出循环
//
//			if i >= j {
//
//				break
//
//			}
//
//			// 交换左右两侧的值
//
//			values[i], values[j] = values[j], values[i]
//
//		}
//
//		// 将分水岭移到哨兵相遇点
//
//		values[left] = values[i]
//
//		values[i] = temp
//
//		// 递归，左右两侧分别排序
//
//		quickSort(values, left, i-1)
//
//		quickSort(values, i+1, right)
//
//	}
//
//}
//
//func QuickSort(values []int) {
//
//	quickSort(values, 0, len(values)-1)
//
//}
//
//func main() {
//	arr:=[]int{1,2,3,4,5,6,7}
//	fmt.Println(arr)
//
//}
