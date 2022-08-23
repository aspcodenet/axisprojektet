package main

import "fmt"

func main() {
	for {
		var val int
		fmt.Println("1. Skapa device")
		fmt.Println("2. Lista alla")
		fmt.Println("3. Ändra device")
		fmt.Println("4. Sök")
		fmt.Println("5. Avsluta")

		fmt.Printf("Välj:")

		_, e := fmt.Scanln(&val)
		if e != nil {
			fmt.Println("Enter a number please")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if val < 1 || val > 5 {
			fmt.Println("Enter a number between 1 and 5 please")
			continue
		}

		if val == 5 {
			break
		}

		fmt.Printf("Du valde: %d", val)

	}

}

// ag får följande när jag kör go mod init main

// go: modules disabled by GO111MODULE=off; see 'go help modules'
