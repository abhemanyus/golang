package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, want, got Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("want %s, got %s\n", want, got)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{0}
		wallet.Deposit(20)
		got := wallet.Balance()
		want := Bitcoin(20.0)
		assertBalance(t, want, got)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{20}
		wallet.Withdraw(10)
		got := wallet.Balance()
		want := Bitcoin(10)
		assertBalance(t, want, got)
	})

	t.Run("overdraw", func(t *testing.T) {
		wallet := Wallet{10}
		err := wallet.Withdraw(20)
		want := Bitcoin(10)
		got := wallet.Balance()

		assertError(t, err, "cannot withdraw, insufficient funds")

		assertBalance(t, want, got)
	})

}
