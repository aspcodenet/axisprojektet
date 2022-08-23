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

/*
[]string{"",""} funkar för mig. Kanske nytt i go?
*/

func main() {

	arr1 := [5]int{} //not initialized

	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr1)
	arr1[1] = 12334
	fmt.Println(arr2)
	fmt.Println(arr3)

	var cars = []string{"Volvo", "BMW", "Ford", "Mazda", "Renault"}
	fmt.Println(cars[0])
	fmt.Println(cars[1])

	cars[3] = "Rolls Royce"

	for index := 0; index < len(cars); index++ {
		x := cars[index]
		fmt.Println(x)
	}

	for index, carName := range cars {
		if index == 0 {
			fmt.Println("Detta är den första bilen")
		}
		fmt.Println(carName)
	}

	for _, carName := range cars {
		fmt.Println(carName)
	}

	// var arr1 = [3]int{1,2,3}

	// var spelare = [111]string{"Zlatan", "Fredrik Ljungberg"}

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
