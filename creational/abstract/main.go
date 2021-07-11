package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShort := nikeFactory.makeShort()
	adidasShoe := adidasFactory.makeShoe()
	adidasShort := adidasFactory.makeShort()

	priintShoeDetail(nikeShoe)
	priintShortDetail(nikeShort)
	priintShoeDetail(adidasShoe)
	priintShortDetail(adidasShort)
}

func priintShoeDetail(s iShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func priintShortDetail(s iShort) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
