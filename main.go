package main

import (
	"fmt"
)

func printMenu() {
	fmt.Println("1. Skapa device")
	fmt.Println("2. Lista alla")
	fmt.Println("3. Ändra device")
	fmt.Println("4. Sök")
	fmt.Println("5. Avsluta")
}

func main() {
	for {
		var val int
		fmt.Printf("Välj:")
		printMenu()
		val = inputIntData(1, 5, "Ange val:")
		switch val {
		case 5:
			break
		case 1:
			createDevice()
		case 2:
			listAll()
		case 3:
			changeDevice()
		case 4:
			search()
		}
		fmt.Printf("Du valde: %d", val)

	}
}

func createDevice() {

}

func listAll() {

}
func changeDevice() {

}
func search() {

}

// ag får följande när jag kör go mod init main

// go: modules disabled by GO111MODULE=off; see 'go help modules'
