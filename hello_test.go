package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Rectangle struct {
	width  float64
	height float64
}

// TestHello to test Hello

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() //serve quando c'e' un errore, il terminal ci dara' la linea giusta
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Simone", "English")
		want := "Hello, Simone"
		assertCorrectMessage(t, got, want)
	})

	t.Run("got an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Simone", "Spanish")
		want := "Hola, Simone"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Simone", "French")
		want := "Bonjour, Simone"
		assertCorrectMessage(t, got, want)
	})
}

func TestSum(t *testing.T) {
	calc := Sum(3, 5)
	res := 8
	if calc != res {
		t.Errorf("got '%d' but want '%d'", calc, res)
	}
}

func ExampleSum() {
	sum := Sum(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestRepetitions(t *testing.T) {
	in := MyRepeat("x", 2)
	out := "xx"
	if in != out {
		t.Errorf("got '%s' but want '%s'", in, out)
	}

}

func TestRepStandard(t *testing.T) {
	in := RepStandard("x", 3)
	out := "xxx"
	if in != out {
		t.Errorf("got '%s' but want '%s'", in, out)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyRepeat("a", 8)
	}
}

func ExampleMyRepeat() {
	r := MyRepeat("a", 5)
	fmt.Println(r)
	// Output: aaaaa
}

func TestSumArray(t *testing.T) {
	t.Run("fixed array of 5 members", func(t *testing.T) {
		numbers := [5]int{2, 4, 6, 8, 1}
		got := SumArray(numbers)
		want := 21
		if want != got {
			t.Errorf("got %d, but want %d given %v", got, want, numbers)
		}
	})

	t.Run("using slices, any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := SumSlice(numbers)
		want := 6
		if want != got {
			t.Errorf("got %d, but want %d given %v", got, want, numbers)
		}
	})
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 20.0}
	got := Perimeter(rectangle)
	want := 60.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 20.0}
	got := Area(rectangle)
	want := 200.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestDeposit(t *testing.T) {

	w := wallet{}

	w.deposit(bitcoin(10))

	got := w.balance()

	//fmt.Println("address of saldo in test function is ", &w.saldo)
	want := bitcoin(10)

	assert.Equal(t, want, got, "Test deposit function failed")
}

func TestWithdrawal(t *testing.T) {

	/*
		w := wallet{saldo: 12}

		w.withdrawal(bitcoin(10))

		got := w.balance()
		want := bitcoin(2)

		assert.Equal(t, want, got, "Withdrawal function incorrect")
	*/

	areatests := []struct {
		amount bitcoin
		w      wallet
		exp    bitcoin
		desc   string
	}{
		{
			10, wallet{0}, 0, "empty wallet should return an empty wallet",
		},
		{
			10, wallet{15}, 5, "enough money in the wallet should return the difference",
		},
		{
			8, wallet{3}, 5, "example of failing test - don't have enough money, but it returns more money than I had",
		},
	}

	for _, el := range areatests {
		el.w.withdrawal(el.amount)
		el.amount = el.w.balance()
		assert.Equal(t, el.exp, el.amount)
	}

}
