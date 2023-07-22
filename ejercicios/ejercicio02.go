package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error
var texto string

func TablaMultiplicar() string {
	for {
		fmt.Printf("Ingrese numero 1 : ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El dato ingresado es incorrecto" + err.Error())
				continue
			}
		}
		break
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d\n", numero1, i, numero1*i)
	}

	return texto
}
