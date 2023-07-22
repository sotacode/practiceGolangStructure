package arreglos_slices

import "fmt"

var tablaSlice []int = []int{2, 5, 4}

var arreglo [10]int = [10]int{6, 6, 3, 98, 54, 23, 12, 3, 4, 6}

func MuestroSlices() {
	fmt.Println(tablaSlice)

	porcion := arreglo[3:]   //Slice creado desde un vector, desde la posicion 3 hasta el final
	porcion2 := arreglo[:4]  //desde el principio hasta la pos 4
	porcion3 := arreglo[2:5] //desde la posicion 2 al 5
	fmt.Println(arreglo)
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	//hacemos una slice con 5 elementos y 20 de capacidad maxima
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d\n", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d\n", len(nums), cap(nums))
}
