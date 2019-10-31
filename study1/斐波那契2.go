package main

func main(){
	feibo(3)
}

//  1  1    2   3   5  8
//  1   2    3   4   5  6
func feibo(n int)int{
	if n ==1 || n==2{
		return 1


	}else {
		//后面的一个  加上  前面的 一个 斐波那契数
		return feibo(n-1)+feibo((n-2))
	}


}