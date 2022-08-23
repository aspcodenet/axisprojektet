package main

import "fmt"

func main() {
	var val int
	fmt.Println(`1. Skapa device
2. Lista alla	
3. Ändra device
4. Sök
5. Avsluta`)

	fmt.Printf("Välj:")
	fmt.Scanln(&val) // ref

	fmt.Printf("Du valde: %d", val)

	// fmt.Println("1. Skapa device")
	// fmt.Println("2. Lista alla")
	// fmt.Println("3. Ändra device")
	// fmt.Println("4. Sök")
	// fmt.Println("5. Avsluta")

}

// ag får följande när jag kör go mod init main

// go: modules disabled by GO111MODULE=off; see 'go help modules'
