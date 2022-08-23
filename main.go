package main

import (
	"fmt"
)

func inputIntData(min int, max int, prompt string) int {
	for {
		fmt.Print(prompt)
		var val int
		_, e := fmt.Scanln(&val)
		if e != nil {
			fmt.Println("Enter a number please")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if val < min || val > max {
			fmt.Println("Enter a number between 1 and 5 please")
			continue
		}
		return val
	}

}

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
