package main

var globalVar = 99

var saldo = 100

func withdraw(belopp int) (int, string) {
	if belopp > 200 {
		return saldo, "Får bara ta ut 200 kr per gång"
	}
	if belopp > saldo {
		return saldo, "Du har inte så mycket pengar"
	}
	saldo -= belopp
	return saldo, ""
}

func add(x int, y int) (total int) {
	total = x + y
	if total > 100 {
		total = total - 100
	}
	return
}

func add1(x int, y int) int {
	total := x + y
	if total > 100 {
		total = total - 100
	}
	return total
}
