package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Test Wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
