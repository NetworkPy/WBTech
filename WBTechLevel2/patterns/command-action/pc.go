package main

import "fmt"

type pc struct {
	isRunning bool
}

func (p *pc) on() {
	p.isRunning = true
	fmt.Println("Turning pc on")
}

func (p *pc) off() {
	p.isRunning = false
	fmt.Println("Turning pc off")
}
