package main

import (
	"fmt"
)

func main() {
	var prethents, rub, cost, max int

	fmt.Scan(&prethents)
	fmt.Scan(&rub)
	mass := make([]int, prethents)
	for i:=0; i<prethents; i++{
		fmt.Scan(&cost)
		mass[i]=cost
	}
	for i:=0; i<=rub; i++{
		rub1:= i
		for _,cost := range mass{
			if rub1>= cost{
				rub1-=cost
			}
		}
		if max<rub1{
			max = rub1
		}
	}
	fmt.Println(max)
}