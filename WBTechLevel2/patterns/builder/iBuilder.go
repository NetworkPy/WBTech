package main

type iBuilder interface {
	setGpuType()
	setCoolingType()
	setDiskType()
	getPc() pc
}

func getBuilder(builderType string) iBuilder {
	switch builderType {
	case "office":
		return &officeBuilder{}
	case "game":
		return &gameBuilder{}
	default:
		return nil
	}
}
