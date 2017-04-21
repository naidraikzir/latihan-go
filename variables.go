package main

import "fmt"

func main() {
	fmt.Println("# Declaration")
	declaration()
	fmt.Println()

	fmt.Println("# Assignment")
	assignment()
	fmt.Println()

	fmt.Println("# Declaration with initial value")
	declareWithInitial()
	fmt.Println()

	fmt.Println("# Type inference")
	typeInference()
	fmt.Println()

	fmt.Println("# Multiple declaration")
	multipleDeclaration()
	fmt.Println()

	fmt.Println("# Multiple declaration with different types")
	multipleDeclarationDifferentTypes()
	fmt.Println()

	fmt.Println("# Shorthand declaration")
	shorthand()
	fmt.Println()
}

func declaration() {
	var age int
	fmt.Println("My age is", age)
}

func assignment() {
	var age int
	fmt.Println("My age is", age)
	age = 27
	fmt.Println("My age is", age)
	age = 54
	fmt.Println("My new age is", age)
}

func declareWithInitial() {
	var age int = 27
	fmt.Println("My age is", age)
}

func typeInference() {
	var age = 27
	fmt.Println("My age is", age)
}

func multipleDeclaration() {
	var width, height int = 100, 50
	fmt.Println("Width is", width, "Height is", height)
}

func multipleDeclarationDifferentTypes() {
	var (
		name   = "naidraikzir"
		age    = 27
		height int
	)
	fmt.Println("My name is", name, ", I am", age, ", my height is", height)
}

func shorthand() {
	name, age := "naidraikzir", 27
	fmt.Println("My name is", name, ", I am", age)
}
