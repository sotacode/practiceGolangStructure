package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error

func GetNumber() {
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
}
