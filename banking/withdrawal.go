package banking

import "errors"

var TooMuchPerWithdrawl = errors.New("Max 2000 per uttag")
var UnsifficientSaldo = errors.New("Du har inte så mycket ")

func Withdraw(kontoNummer string, belopp int) (int, error) {
	currentSaldo := getSaldoFor(kontoNummer)
	if belopp > 2000 {
		return currentSaldo, TooMuchPerWithdrawl // Max 2000 kr per uttag
	}
	if belopp > currentSaldo {
		return currentSaldo, UnsifficientSaldo //Har ju inte tilltäckligt på kontot
	}
	currentSaldo -= belopp
	//Save ny saldo
	return currentSaldo, nil // Detta gick bra!

}

func getSaldoFor(kontoNummer string) int {
	return 200
}
