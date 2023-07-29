package main

import "strings"
import "fmt"

func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}


func RepeatStrings(character string, times int) string {
	return strings.Repeat(character, times)
}

func main(){
	fmt.Println(strings.Repeat("a", 5))
}
