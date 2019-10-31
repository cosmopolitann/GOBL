package main

import "fmt"

func main(){
	arr:=[7]int{12,44,32,68,50,90,35}

	for i:=0;i<len(&arr)-1 ;i++  {
		for j:=0;j<len(arr)-1-i;j++  {
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}


	}

fmt.Println(arr)

}
