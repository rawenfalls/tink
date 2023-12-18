package main

import (
	"fmt"
)

func main() {
	var n, m, q, enter, enter2 int
	var str string
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&q)
	mass := make([]int, n)
	myMap := make(map[int][]int)
	for i:=0;i<n;i++{
		fmt.Scan(&enter)
		mass[i]=enter
	}
	for i:=0;i<m;i++{
		fmt.Scan(&enter)
		fmt.Scan(&enter2)
		myMap[enter] = append(myMap[enter], enter2)
		myMap[enter2] = append(myMap[enter2], enter)
	}
	for i:=0;i<q;i++{
		fmt.Scanf("%s %d", &str, &enter)
		if str[0] == '?'{
			fmt.Println(mass[enter-1])
		}
		if str[0] == '+'{
			fmt.Scan(&enter2)
			for _, i := range myMap[enter]{
				mass[i-1]+=enter2
			}
		}
	}

}