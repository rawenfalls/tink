package main

import (
	"fmt"
)

func main() {
	var n, q, enter, enter2, enter3, enter4, max, p int
	var str string
	fmt.Scan(&n)
	fmt.Scan(&q)
	mass := make([]int, n)
	for i:=0;i<n;i++{
		fmt.Scan(&enter)
		mass[i]=enter
	}
	for i:=0;i<q;i++{
		fmt.Scanf("%s %d %d %d", &str, &enter, &enter2, &enter3)
		if str[0] == '?'{
			fmt.Scan(&enter4)
			max =0 
			p =0 
			for i:=enter; i<=enter2; i++{
				if (mass[i] > max){
					max = mass[i]
					p=i
				}
			}
			fmt.Println(enter3*p+enter4)
		}
		if str[0] == '+'{
			for i:=enter; i<=enter2; i++{
				mass[i]+=enter3
			}
		}
	}

}