package main

import (
	"fmt"
	"strings"
)

const spanish = "Spanish"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const french = "French"

// Hello function
func Hello(name string, language string) string {

	if name == "" {
		return "Hello, world"
	}

	prefix := helloPrefix
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	return prefix + name
}

// Sum that calculates the sum
func Sum(a, b int) int {

	return a + b
}

//MyRepeat a char n times
func MyRepeat(c string, n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += c
	}
	return s
}

//RepStandard uses the Standard Library
func RepStandard(c string, n int) string {
	return strings.Repeat(c, n)
}

//SumArray sums the numbers in the array
func SumArray(n [5]int) int {
	sum := 0
	for _, n := range n {
		sum += n

	}
	return sum
}

//SumSlice sums the numbers in the slice
func SumSlice(s []int) int {
	sum := 0
	for _, s := range s {
		sum += s
	}
	return sum
}

//Perimeter calculates rectangle's perimeter
func Perimeter(r Rectangle) float64 {
	return (r.width + r.height) * 2
}

//Area calculates rectangle's area
func Area(r Rectangle) float64 {
	return r.width * r.height
}

type bitcoin int

type wallet struct {
	saldo bitcoin
}

func (w *wallet) deposit(amount bitcoin) {
	//fmt.Println("address of saldo in deposit function is ", &w.saldo)
	w.saldo += amount
}

func (w *wallet) balance() bitcoin {
	return w.saldo

}

func (w *wallet) withdrawal(amount bitcoin) {
	if amount > w.saldo {
		fmt.Println("don't have enough bitcoins")
	} else {
		w.saldo -= amount
	}
}

func main() {
	s := "porcodio"
	fmt.Println(Hello(s, "Italian"))

}
