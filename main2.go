package main

import (
	"fmt"
)

func main() {
	var num, workers, w, sum int

	fmt.Scanln(&num)
	for i:=0; i<num; i++{
		fmt.Scanln(&workers)
		//mass := make([]int, workers)
		sum = 0
		for j:=0; j<workers; j++{
			fmt.Scan(&w)
			//mass[j]=w
			sum+=w
		}
		if sum>=(2*(workers-1)){
			fmt.Println("Yes")
		}else{
			fmt.Println("No")
		}
	}
}