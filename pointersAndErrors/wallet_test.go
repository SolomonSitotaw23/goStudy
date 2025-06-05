package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10.0)

		assertBalance(t, want, wallet)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20.0),
		}
		err := wallet.Withdraw(Bitcoin(10.0))

		assertNoError(t, err)

		want := Bitcoin(10.0)

		assertBalance(t, want, wallet)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)

		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, startingBalance, wallet)

	})
}

func assertBalance(t testing.TB, want Bitcoin, wallet Wallet) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("expected %s but got %s", want, got)
	}
}

// assertError checks that an error was returned and that its message matches the expected string.
// It fails the test if no error is returned or if the error message is incorrect.
func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got an error but didn't want one")
	}

}
