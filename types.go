package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println()
	fmt.Println("# Boolean")
	boolean()
	fmt.Println()

	fmt.Println("# Signed integer")
	signedInteger()
	fmt.Println()

	fmt.Println("# Type and size")
	typeAndSize()
	fmt.Println()

	fmt.Println("# Floating point size")
	floatingPointSize()
	fmt.Println()

	fmt.Println("# Complex types")
	complexTypes()
	fmt.Println()

	fmt.Println("# Strings")
	strings()
	fmt.Println()

	fmt.Println("# Type conversion")
	typeConversion()
	fmt.Println()

	fmt.Println("# Assign and convert")
	assignAndConvert()
	fmt.Println()
}

func boolean() {
	var a bool = true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c, "👉  a && b")
	d := a || b
	fmt.Println("d:", d, "👉  a || b")
}

func signedInteger() {
	var a int = 89
	b := 95
	fmt.Println("Value of a is", a, "and b is", b)
}

func typeAndSize() {
	var a int = 89
	b := 95
	fmt.Println("Value of a is", a, "and b is", b)
	fmt.Printf("Type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\nType of b is %T, size of b is %d\n", b, unsafe.Sizeof(b))
}

func floatingPointSize() {
	a, b := 5.67, 8.97
	fmt.Println("a:", a, "b:", b)
	fmt.Printf("Type of a %T and b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("Sum", sum, "Diff", diff)

	no1, no2 := 56, 89
	fmt.Println("\nno1:", no1, "no2:", no2)
	fmt.Println("Sum", no1+no2, "Diff", no1-no2)
}

func complexTypes() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	fmt.Println("c1:", c1, "c2:", c2)
	cadd := c1 + c2
	fmt.Println("Sum:", cadd)
	cmul := c1 * c2
	fmt.Println("Multiply:", cmul)
}

func strings() {
	first := "Rizki"
	last := "Ardian"
	fmt.Println("first:", first, ", last:", last)
	name := first + " " + last
	fmt.Println("My name is", name)
}

func typeConversion() {
	i := 55
	j := 67.8
	fmt.Println("i:", i, ", j:", j)
	sum := i + int(j)
	fmt.Println("Sum:", sum)
}

func assignAndConvert() {
	i := 10
	var j float64 = float64(i)
	fmt.Println("j:", j)
}
