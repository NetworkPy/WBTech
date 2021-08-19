package main

type Subaru struct {
	wheelDiameter int
}

func (t *Subaru) accept(v visitor) {
	v.visitForSubaru(t)
}

func (t *Subaru) getType() string {
	return "subaru"
}
