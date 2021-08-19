package main

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildPc() pc {
	d.builder.setGpuType()
	d.builder.setCoolingType()
	d.builder.setDiskType()
	return d.builder.getPc()
}
