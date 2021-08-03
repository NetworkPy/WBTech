package main

import "fmt"

var justString string // Global virable

// We can see that justString is global. Another developers get our code and work with it.
// If we can 100500+ link for justString that another person have no chance to see all links... It's impossible and hard! Don't use it beacuse is hard to support.
// Local virable is the best. We can use return in function and change value for justString.
// Global virable support work with configs. We read configuration file and use value from it in all program.


func someFunc() {
	v := createHugeString()
	justString = v[:100]
}

func main() {
	someFunc()
	fmt.Println(justString)
}

func createHugeString() string {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "a"
	}
	return str
}
