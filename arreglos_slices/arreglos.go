package arreglos_slices

import "fmt"

var tabla [10]int
var tabla2 [10]int = [10]int{10, 0, 59}

var matriz [10][10]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54
	fmt.Println(tabla)
	fmt.Println(tabla2)

	tabla3 := [10]string{"Nelsota"}
	fmt.Println(tabla3)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][9] = 999
	fmt.Println(matriz)
}
