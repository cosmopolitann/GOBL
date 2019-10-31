package main

import "fmt"

func main(){
	//吃桃子

	//思路分析

	/*
	n           (n+1)*2
	7           22
	8           10
	9           4
	10          1  taozi
	 */





}
func  testtao(a int)int{
	if a >10 || a<1{

		fmt.Println("输入的天数不正确")
	}
	if a==10{
		return 1
	}else{


		return  (testtao(a+1)+1)*2
	}



}