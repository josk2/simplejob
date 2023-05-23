package tddexample

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertErr(t, err, ErrNotEnoughFund)
		assertBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertErr(t, err, ErrNotEnoughFund)
		assertBalance(t, &wallet, Bitcoin(20))
	})
}

func assertBalance(t *testing.T, w *Wallet, want Bitcoin) {
	t.Helper()
	if w.Balance() != want {
		t.Errorf("got %#v want %#v", w.Balance(), want)
	}
}

func assertErr(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
