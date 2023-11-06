package ptr

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("Wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	}
	assertBalance := func(t testing.TB, wallet Wallet, want BitCoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, BitCoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}
		wallet.Withdraw(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := BitCoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(BitCoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}
