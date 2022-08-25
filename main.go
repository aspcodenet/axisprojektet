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

// // OOP ???
// // hmmm... lets start with user defined types
// // HockeyPlayer - name, age, jersey
// type HockeyPlayer struct {
// 	Name         string
// 	Age          int
// 	JerseyNumber int
// 	PlayedGames  []Games
// }
// type Games struct {
// }

// func createPlayer() *HockeyPlayer {
// 	var b = HockeyPlayer{
// 		Name:         "tefsan",
// 		Age:          12,
// 		JerseyNumber: 123,
// 	};
// 	var c = HockeyPlayer{
// 		Name:         "tefsan",
// 		Age:          12,
// 		JerseyNumber: 123,
// 	};
// 	return &c;

// 	return &HockeyPlayer{
// 		Name:         "tefsan",
// 		Age:          12,
// 		JerseyNumber: 123,
// 	}
// }
// var p HockeyPlayer
// p.Name = "Stefan"
// p.Age = 123

// var players []HockeyPlayer
// players = append(players, p)

// //var bestPlayer *HockeyPlayer
//
//	p2 := HockeyPlayer{
//		Name:         "tefsan",
//		Age:          12,
//		JerseyNumber: 123,
//	}
type Device struct {
	Name          string
	Color         string
	Manufacturer  string
	WeightInGrams int
	Test          map[string]interface{}
}

func main() {
	var aa map[string]interface{}
	aa = make(map[string]interface{})
	aa["Hello"] = 12
	aa["Hello2"] = 13

	var a int
	a = 1
	if a == 0 || a == 1 ||
		a == 2 {
		fmt.Println("sdadsasda")

	}

	//	stackheap.Demo()
	var devices []Device
	for {
		var val int
		fmt.Printf("Välj:")
		printMenu()
		val = inputIntData(1, 5, "Ange val:")
		switch val {
		case 5:
			break
		case 1:
			newDevice := createDevice()
			devices = append(devices, *newDevice)
		case 2:
			listAll(&devices)
		case 3:
			changeDevice(&devices)
		case 4:
			search(&devices)
		}
		fmt.Printf("Du valde: %d", val)

	}
}

func createDevice() *Device {
	var device Device

	fmt.Println("*** Ny device ***")
	fmt.Print("Namn:")
	fmt.Scanln(&device.Name)
	fmt.Print("Manufacturer:")
	fmt.Scanln(&device.Manufacturer)
	fmt.Print("Color:")
	fmt.Scanln(&device.Color)
	fmt.Print("Weight:")
	fmt.Scanln(&device.WeightInGrams)

	fmt.Println("*** Device added ***")
	return &device

}

func listAll(devices *[]Device) {
	fmt.Println("*** All our devices ***")
	for _, device := range *devices {
		fmt.Println(device)
	}

}
func changeDevice(devices *[]Device) {
	fmt.Println("*** Select one to modify ***")
	for index, device := range *devices {
		fmt.Printf("%d %s\n", index+1, device.Name)
	}
	sel := inputIntData(1, len(*devices)+1, "Ange vilken")
	var name, manufacturer, color string
	var weightInGrams int
	fmt.Print("Namn:")
	fmt.Scanln(&name)
	fmt.Print("Manufacturer:")
	fmt.Scanln(&manufacturer)
	fmt.Print("Color:")
	fmt.Scanln(&color)
	fmt.Print("Weight:")
	fmt.Scanln(&weightInGrams)

	(*devices)[sel].Name = name
	(*devices)[sel].Manufacturer = manufacturer
	(*devices)[sel].Color = color
	(*devices)[sel].WeightInGrams = weightInGrams

}
func search(devices *[]Device) {

}

// ag får följande när jag kör go mod init main

// go: modules disabled by GO111MODULE=off; see 'go help modules'
