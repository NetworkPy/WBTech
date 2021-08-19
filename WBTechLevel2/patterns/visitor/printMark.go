package main

import (
	"fmt"
)

type Printer struct {
}

func (p *Printer) visitForSubaru(s *Subaru) {
	fmt.Println(s.wheelDiameter)
}

func (p *Printer) visitForBmw(s *Bmw) {
	fmt.Println(s.wheelDiameter)
}

func (p *Printer) visitForAudi(s *Audi) {
	fmt.Println(s.wheelDiameter)
}
