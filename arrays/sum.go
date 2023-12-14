package arrays

type Account struct {
	Name    string
	Balance int
}

type Transaction struct {
	From string
	To   string
	Sum  int
}

func NewTransaction(from, to Account, amount int) Transaction {
	return Transaction{
		From: from.Name,
		To:   to.Name,
		Sum:  amount,
	}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, account, applyTransaction)
}

func applyTransaction(account Account, transaction Transaction) Account {
	if account.Name == transaction.From {
		account.Balance -= transaction.Sum
	}
	if account.Name == transaction.To {
		account.Balance += transaction.Sum
	}

	return account
}

func Find[T any](haystack []T, magnet func(T) bool) (needle T, found bool) {
	for _, item := range haystack {
		if magnet(item) {
			return item, true
		}
	}

	return
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

func Reduce[T, U any](items []T, initial U, op func(U, T) U) U {
	result := initial
	for _, item := range items {
		result = op(result, item)
	}
	return result
}
