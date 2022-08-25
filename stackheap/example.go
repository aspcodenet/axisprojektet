package stackheap

import "fmt"

func change1(varde int) { // COPY BY VALUE
	varde = varde + 1
	fmt.Println(varde) // -> 4
}

func demo1() { // Skolfråga: vad blir utskriften
	var varde = 3
	fmt.Println(varde)
	change1(varde)
	fmt.Println(varde) // -> 3
}

// REF
func change2(varde *int) { // COPY BY VALUE
	*varde = *varde + 1
	fmt.Println(varde) // -> 4
}

func demo2() { // Skolfråga: vad blir utskriften
	var varde = 3
	fmt.Println(varde) // -> 3
	change2(&varde)
	fmt.Println(varde) // -> 4
}

func change3(arrayen [3]string) { // COPY BY VALUE
	arrayen[0] = "Pelle"
	fmt.Println(arrayen)
}

func demo3() { // Skolfråga: vad blir utskriften
	var arrayen = [3]string{"Stefan", "Lisa", "Nisse"}
	fmt.Println(arrayen) // -> 3
	change3(arrayen)
	fmt.Println(arrayen) // -> 4
}

func change4(arrayen []string) { // COPY BY VALUE
	arrayen[0] = "Pelle"
	fmt.Println(arrayen)
}

func demo4() { // Skolfråga: vad blir utskriften
	var arrayen = []string{"Stefan", "Lisa", "Nisse"}
	fmt.Println(arrayen) // -> 3
	change4(arrayen)
	fmt.Println(arrayen) // -> 4
}

func change5(arrayen []string) { // COPY BY VALUE
	arrayen[0] = "Pelle"
	arrayen = append(arrayen, "Arne")
	fmt.Println(arrayen)
}

func demo5() { // Skolfråga: vad blir utskriften
	var arrayen = []string{"Stefan", "Lisa", "Nisse"}
	fmt.Println(arrayen)
	change5(arrayen)
	fmt.Println(arrayen)
}

//	Array ["Stefan"]["Lisa"]["Nisse"]
//
// slice = struct  3 fields cap, len 3, ptr -> ["Stefan"]["Lisa"]["Nisse"]["Arne"]
func Demo() {
	demo5()
}
