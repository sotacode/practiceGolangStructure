package iteraciones

import "fmt"

func Iterar() {
	i := 0
	//Similes al while
	for {
		//itera eternamente hasta un break
		fmt.Println("Iterando")
		break
	}
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 50; j += 5 {
		fmt.Println(j)
	}
}
