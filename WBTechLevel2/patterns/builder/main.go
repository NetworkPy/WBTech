package main

import "fmt"

func main() {
	gameBuilder := getBuilder("game")
	officeBuilder := getBuilder("office")
	director := newDirector(gameBuilder)
	
	gamePc := director.buildPc()
	director.setBuilder(officeBuilder)
	officePc := director.buildPc()

	fmt.Printf("Office PC Cooling Type: %s\n", officePc.gpuType)
	fmt.Printf("Office PC Cooling Type: %s\n", officePc.coolingType)
	fmt.Printf("Office PC Cooling Type: %s\n", officePc.diskType)

	fmt.Printf("\nGame PC GPU Type: %s\n", gamePc.gpuType)
	fmt.Printf("Game PC Cooling Type: %s\n", gamePc.coolingType)
	fmt.Printf("Game PC Disk Type: %s\n", gamePc.diskType)

}
