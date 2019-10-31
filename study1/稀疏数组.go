
package main

import "fmt"

type Node struct {
	row int
	col int
	val int
}

func main() {

	//程序就是数据结构  加算法

	//稀疏数组

	var arr [10][10]int
	arr[1][3]=1
	arr[2][4]=2

	for _,v:=range arr{

		for _,v:=range v{
			fmt.Print(v)
		}
		fmt.Println()
	}
	var s Node=Node{
		row: 11,
		col: 11,
		val: 0,
	}
	var  xishu []Node
	xishu=append(xishu,s)

	for i,v:=range arr{
		for j,v:=range v{

			if v!=0{
				a:=Node{
					row: i,
					col: j,
					val: v,
				}
				xishu=append(xishu ,a)
			}
		}
	}

	for k,v:=range xishu{
		fmt.Println(k,v)


	}







}
