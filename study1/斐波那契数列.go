package main

import "fmt"

func main(){
	//  1   1    2    3     5   8

	s:=testfeibo(5)
	fmt.Println(s)
	//sadfa



}
func testfeibo(n int)int{
	if n==1 || n==2{
		return 1


	}else{
		return testfeibo(n-1)+testfeibo(n-2)

	}

}
