package main

import "fmt"

type Fruit int
type Animal int

const (
	Apple Fruit = iota
	Orange
	Banana
)

const (
	Monkey Animal = iota
	Elephant
	Pig
)

func main() {
	fmt.Println("こんにちは せかい")

	var fruit Fruit = Apple
	fmt.Println(fruit)

	//fruit = Elephant // Fruit型のfruitにAnimal型のElephantを代入するとエラー
	fruit = Banana
	fmt.Println(fruit)
}
