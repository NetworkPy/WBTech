package main

type Audi struct {
	wheelDiameter int
}

func (t *Audi) accept(v visitor) {
	v.visitForAudi(t)
}

func (t *Audi) getType() string {
	return "Audi"
}
