package main

import "fmt"

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
