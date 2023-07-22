package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"src/github.com/sotacode/practiceGolangStructure/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	//Borra archivo y lo crea nuevamente en caso de existir archivo
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo: ", err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido: ")
		return
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al leer el archivo: ", err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error al escribir en el archivo: ", err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer archivo: " + err.Error())
		return
	}
	fmt.Println(string(archivo))

	archivo2, err2 := os.Open(fileName)
	if err2 != nil {
		fmt.Println("Error al leer archivo: " + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo2)
	//recorremos linea por linea
	for scanner.Scan() {
		fmt.Println("> " + string(scanner.Text()))
	}
	archivo2.Close()
}
