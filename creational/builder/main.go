package main

import (
	"fmt"
)

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num floor Type: %d\n", normalHouse.floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("iglooHouse House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("iglooHouse House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("iglooHouse House Num floor Type: %d\n", iglooHouse.floor)
}
