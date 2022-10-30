package main

import "fmt"

func helloWorld() {
	const (
		a = iota
	)
	// var b bool
	// What is an iota?
	// Iota is a constant generator which generates a sequence of integers starting with 0.
	// It is used in place of the numbers when we don't want to write them explicitly.
	// It can be used in place of numbers in const, var, for loop, switch statements.
	// It can be used in place of numbers in array, slice, map, struct, interface.
	// It can be used in place of numbers in function calls.
	// It can be used in place of numbers in type declarations.
	// It can be used in place of numbers in type assertions.
	// It can be used in place of numbers in type conversions.
	// It can be used in place of numbers in type switches.
	// It can be used in place of numbers in type specifications.
	// It can be used in place of numbers in type definitions.
	// It can be used in place of numbers in type instantiations.
	// It can be used in place of numbers in type specifications.

	// Write an expression using iota which generates the following sequence of numbers: 10, 20, 30
	const (
		c = iota * 10
		d = iota * 10
		e = iota * 10
		f = iota * 10
	)

	fmt.Println(a, c, d, e, f)

	b := true

	// What does the code below print?
	fmt.Println("Hello, World!", a, b)
	fmt.Println("Hello, World!", a, b)

	s := "hello"
	char := []rune(s)
	char[0] = 'c'

	fmt.Println(char, s)
}

func complexNumbers() {
	// complex numbers
	var c complex64 = 5 + 5i
	fmt.Println(c)
}

func controlStructures() {
	// if statement
	if true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// If elseif
	if false {
		fmt.Println("false")
	} else if true {
		fmt.Println("true")
	} else {
		fmt.Println("false-e")
	}

	// switch statement
	switch "docker" {
	case "linux":
		fmt.Println("linux")
	case "docker":
		fmt.Println("docker")
	default:
		fmt.Println("default")
	}

	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Do while loop
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	// Range
	s := []int{1, 2, 3}
	for i, v := range s {
		fmt.Println(i, v)
	}

	// Continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// For each
	s = []int{1, 2, 3}
	for _, v := range s {
		fmt.Println(v)
	}

	// Nested for loop
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
		}
	}

	// Nested loop with break
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
			if j == 2 {
				break
			}
		}
	}
}

func arrays() {
	// Arrays
	var a [5]int
	a[2] = 7
	fmt.Println(a)

	// Array literal
	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	// Array length
	fmt.Println(len(b))

	// Array of strings
	c := [2]string{"hello", "world"}
	fmt.Println(c[0], c[1])

	// Two dimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Slices
	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d", i, s[i])
	}

	// Slices are like references to arrays
	fmt.Println("s[1:4] ==", s[1:4])

	// missing low index implies 0
	fmt.Println("s[:3] ==", s[:3])

	// missing high index implies len(s)
	fmt.Println("s[4:] ==", s[4:])

	// Slices can contain any type, including other slices
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slice literals
	// A slice literal is like an array literal without the length.
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// Slice defaults
	// When slicing, you may omit the high or low bounds to use their defaults instead.
	// The default is zero for the low bound and the length of the slice for the high bound.
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// Slice length and capacity
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	// You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
	// Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.
	s = s[:0]
	fmt.Println(s, len(s), cap(s))
}

func arrays2() {
	a := [5]int{1, 2, 3, 4, 5}

	s1 := a[2:4]
	s2 := a[1:5]

	// Print the slices with their size and capacity
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	// Create new slice using new keyword
	s3 := new([5]int)[2:4]

	// Print the slices with their size and capacity
	fmt.Println(s3, len(s3), cap(s3))
}

func arraysCopy() {
	// Copy slices
	a := []int{1, 2, 3}
	b := make([]int, 2)
	copy(b, a)
	fmt.Println(b)

	// Append slices
	a = []int{1, 2, 3}
	b = []int{4, 5}
	c := append(a, b...)
	fmt.Println(c)
}

func maps() {
	// Maps
	// A map maps keys to values.
	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	// The make function returns a map of the given type, initialized and ready for use.
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	fmt.Println("map:", m)

	// Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	// Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Exists
	_, prs2 := m["k1"]
	fmt.Println("prs:", prs2)

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// Maps are reference types
	// When maps are passed to a function, the function receives a copy of the reference to the map, so any changes the function makes will be visible through the caller’s map reference.
	// To avoid confusion, it’s a good idea to pass a map as a pointer to a function when the function needs to modify the map.

	// Map literals
	// Map literals are like struct literals, but the keys are required.
	// The following code creates a map that associates strings to ints.
	// The map is initialized with two key/value pairs.
	// The first key/value pair is "a" -> 1 and the second is "b" -> 2.

	// Map literals continued
	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	// The key type can be omitted from the elements of the literal if it’s a string.
	// The value type can be omitted from the elements of the literal if it can be inferred from the map type.

	// Looping over maps example
	// To iterate over key/value pairs in a map, use a range loop.
	// The range form of the loop iterates over a map yielding a pair of values for each iteration.
	// The first value is the key, and the second is the value associated with the key.
	// The order of iteration is not specified and is not guaranteed to be the same from one iteration to the next.
	// If you only want the keys, you can omit the second variable in the range clause.
	// If you only want the values, you can use the blank identifier _ as the key.
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}

	for k, v := range monthdays {
		fmt.Printf("%s has %d days\n", k, v)
	}
}

func main() {
	helloWorld()
	complexNumbers()
	controlStructures()
	arrays()
	arrays2()
	arraysCopy()
	maps()
}
