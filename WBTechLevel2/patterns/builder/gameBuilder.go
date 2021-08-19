package main

type gameBuilder struct {
	coolingType string
	diskType    string
	gpuType     string
}

func newGameBuilder() *gameBuilder {
	return &gameBuilder{}
}

func (g *gameBuilder) setCoolingType() {
	g.coolingType = "Water cooling"
}

func (g *gameBuilder) setDiskType() {
	g.diskType = "SSD"
}

func (g *gameBuilder) setGpuType() {
	g.gpuType = "Discrete"
}

func (g *gameBuilder) getPc() pc {
	return pc{
		coolingType: g.coolingType,
		diskType:    g.diskType,
		gpuType:     g.gpuType,
	}
}
