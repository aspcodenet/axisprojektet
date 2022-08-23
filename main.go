package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Vad heter du?")
	fmt.Scanln(&name)
	fmt.Print("Hur gammal är du?")
	fmt.Scanln(&age)

	fmt.Printf("Hej %s du är %d år\n", name, age)

}

// ag får följande när jag kör go mod init main

// go: modules disabled by GO111MODULE=off; see 'go help modules'
