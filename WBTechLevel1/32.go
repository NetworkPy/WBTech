package main

import (
	"time"
)

// Написать собственную функцию Sleep.

func main() {
	mySleep(1)
}

func mySleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}
