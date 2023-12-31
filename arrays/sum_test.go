package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTailes(t *testing.T) {
	t.Run("Sum tail of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("Safely handle single element slices", func(t *testing.T) {
		got := SumAllTails([]int{1})
		want := []int{0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
	t.Run("Safely handle empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("Mulitply", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3}, 1, multiply), 6)
	})
	t.Run("Concatenate", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}
		AssertEqual(t, Reduce([]string{"a", "b", "c"}, "", concatenate), "abc")
	})
}

func TestBadBank(t *testing.T) {
	riya := Account{Name: "Riya", Balance: 100}
	chris := Account{Name: "Chris", Balance: 75}
	adil := Account{Name: "Adil", Balance: 200}
	transactions := []Transaction{
		NewTransaction(chris, riya, 100),
		NewTransaction(adil, chris, 25),
	}

	newBalanceFor := func(account Account) int {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)
}

func TestFind(t *testing.T) {
	t.Run("find first even", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		firstEven, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		AssertTrue(t, found)
		AssertEqual(t, firstEven, 2)
	})
}

func AssertTrue(t *testing.T, predicate bool) {
	t.Helper()
	if !predicate {
		t.Errorf("Not true")
	}
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("Didn't want %v", want)
	}
}
