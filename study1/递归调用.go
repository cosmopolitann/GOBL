package main

import "fmt"
//栈
//一个函数 在函数体内调用了本身
func main(){

	test(4)


}
func test(a int){

	if a>2{

		a--
		test(a)
fmt.Println(a)

	}
	fmt.Println(a)

}