package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Zenny(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Zenny(20)}
		err := wallet.Withdraw(Zenny(10))
		assertBalance(t, wallet, Zenny(10))
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Zenny(5)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Zenny(10))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Zenny) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error, didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
