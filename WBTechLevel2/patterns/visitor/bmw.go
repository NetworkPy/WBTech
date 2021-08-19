package main

type Bmw struct {
	wheelDiameter int
}

func (t *Bmw) accept(v visitor) {
	v.visitForBmw(t)
}

func (t *Bmw) getType() string {
	return "bmw"
}
