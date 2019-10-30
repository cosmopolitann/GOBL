package main

import "fmt"

//这个是包 名 不一定和 文件名一样  但是一样是为了 方便调用

//函数的调用机制
func main() {
  s:=jisuan(1,2)
  fmt.Println("这是计算后的数字",s)


}

func jisuan(a, b int) int {

	s := a + b
	return s

}
