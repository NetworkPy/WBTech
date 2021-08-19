package main

type officeBuilder struct {
	coolingType string
	diskType    string
	gpuType     string
}

func newOfficeBuilder() *officeBuilder {
	return &officeBuilder{}
}

func (o *officeBuilder) setCoolingType() {
	o.coolingType = "Air cooling"
}

func (o *officeBuilder) setDiskType() {
	o.diskType = "Hard"
}

func (o *officeBuilder) setGpuType() {
	o.gpuType = "Integrated"
}

func (o *officeBuilder) getPc() pc {
	return pc{
		coolingType: o.coolingType,
		diskType:    o.diskType,
		gpuType:     o.gpuType,
	}
}
