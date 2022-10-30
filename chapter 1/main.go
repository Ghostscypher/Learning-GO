package main

import (
	"fmt"
	"strings"
)

func question1() {
	// Create a simple loop with the for construct. Make it loop 10 times and print out
	// the loop counter with the fmt package.
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func question2() {
	// Rewrite the loop from 1. to use goto. The keyword for may not be used.
	i := 0
loop:
	fmt.Println(i)
	i++

	if i < 10 {
		goto loop
	}
}

func question3() {
	// Rewrite the loop again so that it fills an array and then prints that array to the
	// screen.

	var arr [10]int

	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	fmt.Println(arr)
}

func question4() {
	// FizzBuzz[23]
	for i := 1; i <= 23; i++ {
		str := ""

		if i%3 == 0 {
			str += "Fizz"
		}

		if i%5 == 0 {
			str += "Buzz"
		}

		if str == "" {
			str = fmt.Sprintf("%d", i)
		}

		fmt.Println(str)
	}
}

func question5() {
	// 	Create a Go program that prints the following (up to 100 characters):
	// A
	// AA
	// AAA
	// AAAA
	// AAAAA
	// AAAAAA
	// AAAAAAA

	// Crude approach
	for i := 1; i <= 100; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("A")
		}
		fmt.Println()
	}

	// Tidied up version
	for i := 1; i <= 100; i++ {
		fmt.Println(strings.Repeat("A", i))
	}

	// Given solution
	str := "A"
	for i := 1; i <= 100; i++ {
		fmt.Printf("%s\n", str)
		str = str + "A"
	}
}

func question6() {
	// Create a program that counts the number of characters in this string:
	// asSASA ddd dsjkdsjs dk
	// In addition, make it output the number of bytes in that string.
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Println(len(str))
	fmt.Println(len([]byte(str)))
}

func question7() {
	// Extend/change the program from the previous question to replace the three runes
	// at position 4 with ’abc’.
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Println(str[:4] + "abc" + str[7:])
}

func question8() {
	// Write a Go program that reverses a string, so “foobar” is printed as “raboof”.
	str := "foobar"

	// Crude approach
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}

	// Tidied up version
	fmt.Println()
	for i := range str {
		fmt.Print(string(str[len(str)-1-i]))
	}
	fmt.Println()

	// Given solution, uses runes and parallel assignment
	// Not very easy to understand (too wordy and compact) as compared to the above
	// solutions
	reversed := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	fmt.Println(string(reversed))
}

func question9() {
	// Write code to calculate the average of a float64 slice. In a later exercise (Q6 you
	// will make it into a function.
	original := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := original[:]
	var sum, average float64

	for _, v := range slice {
		sum += v
	}

	average = sum / float64(len(slice))
	fmt.Println(average)
}

func main() {
	question1()
	question2()
	question3()
	question4()
	question5()
	question6()
	question7()
	question8()
	question9()
}
