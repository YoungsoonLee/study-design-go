package main

import (
	"fmt"

	"github.com/YoungsoonLee/study-design-go/creational/builder/design"
)

func main() {
	normalBuilder := design.GetBuilder("normal")
	iglooBuilder := design.GetBuilder("igloo")

	director := design.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num floor Type: %s\n", normalHouse.Floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("iglooHouse House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("iglooHouse House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("iglooHouse House Num floor Type: %s\n", iglooHouse.Floor)
}
