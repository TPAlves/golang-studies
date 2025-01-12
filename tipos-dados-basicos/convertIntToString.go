package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func ConvertUseFmt(number int) (text string) {
	return fmt.Sprintf("%d", number)
}

func ConvertUseStrconv(number int) (text string) {
	return strconv.Itoa(number)
}

func main() {
	num := 1000
	fmt.Println(reflect.TypeOf(num))
	fmt.Println((ConvertUseStrconv(num)))
}
