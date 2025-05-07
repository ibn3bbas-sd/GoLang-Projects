package main

func main2() {

	// Variable declaration and initialization
	// Using var keyword
	// var keyword is used to declare a variable in Go.
	// The variable name is followed by the type and then the value.
	//Types of variables in Go:
	// 1. int: Integer type, used for whole numbers.
	// 2. string: String type, used for text.
	// 3. bool: Boolean type, used for true/false values.
	// 4. float32: Floating point type, used for decimal numbers.
	// 5. float64: Double precision floating point type, used for decimal numbers with more precision.
	// 6. complex64: Complex number type, used for complex numbers with float32 real and imaginary parts.
	// 7. complex128: Complex number type, used for complex numbers with float64 real and imaginary parts.
	// 8. byte: Alias for uint8, used for byte values.
	// 9. rune: Alias for int32, used for Unicode code points.
	// 10. uint: Unsigned integer type, used for whole numbers without sign.
	// 11. uint8: Unsigned 8-bit integer type, used for whole numbers without sign.

	var age int = 36
	var name string = "Ibn Abbas"
	var isMarried bool = true
	var height float32 = 172.5

	println(name)
	println(age)
	println(isMarried)
	println(height)
}
