package main


func main(){




}
func test11(n int)int{

	if n==1{
		return 3
	}else{
		return  2*test11(n-1)+1
	}

}