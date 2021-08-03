package main

import (
	"fmt"
	"reflect"
)

// Написать программу, которая в рантайме способна определить тип переменной — int, string, bool, channel из переменной типа interface{}.

func main() {
	channelInt := make(chan int)
	channelStr := make(chan string)
	channelBool := make(chan bool)
	var arr = []interface{}{12, "sam", true, channelInt, channelStr, channelBool}
	for _, val := range arr {
		thirdRuntimeType(val)
	}
}

func firstRuntimeType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	case chan bool:
		fmt.Println("chan bool")
	default:
		fmt.Println("Another type")
	}
}

// TypeOf() from reflect
func secondRuntimeType(x interface{}) {
	val := reflect.TypeOf(x)
	fmt.Println(val)
}

// ValueOf().Type() from reflect
func thirdRuntimeType(x interface{}) {
	val := reflect.ValueOf(x).Type()
	fmt.Println(val)
}
