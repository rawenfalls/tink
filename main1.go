package main

import (
	"fmt"
	"reflect"
)

func chek(word string, letterCountMapKeyWord map[string]int){
	letterCountMap := make(map[string]int)
	for _, char := range word{
		letter := string(char)
		letterCountMap[letter]++
	}
	equal := reflect.DeepEqual(letterCountMap, letterCountMapKeyWord)
	if equal{
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}

func main() {
	var word string
	var num int
	keyWord:= "TINKOFF"

	letterCountMapKeyWord := make(map[string]int)
	for _, char := range keyWord{
		letter := string(char)
		letterCountMapKeyWord[letter]++
	}
	fmt.Scanln(&num)
	for i:=0; i<num; i++{
		fmt.Scanln(&word)
		chek(word,letterCountMapKeyWord)
	}
}