package main

import (
	"fmt"
	"sync"
)

// Чем завершится данная программа?
// Данная программа завершится ошибкой, т.к. в горутины передается на указатель на wg. Получается так,
// что в каждой анонимной функции будет новая копия wg и мы не сможем отследить все Add и Done

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}


// True version for usability
func trueVer() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
