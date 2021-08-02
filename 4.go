package main

import (
	"fmt"
	"runtime"
)

// Реализовать набор из N воркеров, которые читают из канала произвольные данные и выводят в stdout. Данные в канал пишутся из главного потока.
// Необходима возможность выбора кол-во воркеров при старте, а также способ завершения работы всех воркеров.

func main() {
	const numWorker int = 100         // Numbers workers
	channel := make(chan interface{}) // Create channel with type interface for work with different input data

	for i := 0; i < numWorker; i++ {
		go func(num int, c <-chan interface{}) {
			for {
				if value, ok := <-c; ok {
					fmt.Printf("Worker %v value: %v type: %T\n", num, value, value) // Read from channel and print
					runtime.Gosched()                                               // Take moment for work other workers
				}
			}
		}(i, channel)
	}

	strArr := []string{"aafasg", "sgsghs", "skgskg", "sjgwtwtiwt", "skgjjsjagshjng"}
	for i := 0; i < 100; i++ {

		if i == 50 {
			close(channel) // Close channel for end work all workers
			break
		}

		if i < 5 {
			channel <- strArr[i]
			continue
		} else if i < 15 {
			channel <- float32(i)
			continue
		}

		channel <- i
	}
}
