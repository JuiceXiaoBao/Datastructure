package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Valnode struct {
	row int `json:"row"`
	col int `json:"col"`
	val int `json:"val"`
}

func main() {
	//1.先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //篮子

	//2.输出原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//3.转成稀疏数组
	//(1).遍历chessMap，如果我们发现有一个元素的值不为0，创建一个valnode结构体
	//(2).将其放入到对应的切片即可

	var sparseArr []Valnode

	//标准的一个稀疏数组应该还有一个记录元素的二维数组的规模（行和列，默认值）

	valNode := Valnode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个值结点
				valNode := Valnode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	//输出稀疏数组
	fmt.Println("当前的稀疏数组是：")
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	data, err := json.Marshal(sparseArr)
	if err != nil {
		fmt.Println("json marshal err=", err)
	}
	filepath := "d:/chess.txt"
	ioutil.WriteFile(filepath, data, 0666)

	//恢复原始的数组
	//1.打开这个d:/chessmap.data=>恢复原始数组
	//2.这里使用稀疏数组恢复

	//先创建一个原始数组
	var chessMap2 [11][11]int
	var sparseArr2 []Valnode
	//遍历稀疏数组[遍历文件每一行]

	data, err = ioutil.ReadFile(filepath)
	err = json.Unmarshal(data, &sparseArr2)
	if err != nil {
		fmt.Println("json unmarshal err=", err)
	}
	for i, valNode := range sparseArr2 {
		if i != 0 { //跳过第一行的规模表示
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}

	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
