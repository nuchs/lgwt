package arrays

type Transaction struct {
	From string
	To   string
	Sum  int
}

func BalanceFor(transactions []Transaction, name string) int {
	var balance int

	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}

	return balance
}

func Sum(numbers []int) int {
	return Reduce(numbers, 0, func(a, b int) int { return a + b })
}

func SumAll(numbersToSum ...[]int) []int {
	return Reduce(numbersToSum, make([]int, 0), func(a, b []int) []int { return append(a, Sum(b)) })
}

func SumAllTails(numbersToSum ...[]int) []int {
	return Reduce(numbersToSum, make([]int, 0), func(a, b []int) []int {
		if len(b) == 0 {
			return append(a, 0)
		} else {
			tail := b[1:]
			return append(a, Sum(tail))
		}
	})
}

func Reduce[T any](items []T, initial T, op func(T, T) T) T {
	result := initial
	for _, item := range items {
		result = op(result, item)
	}
	return result
}
