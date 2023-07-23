package ejer_interfaces

import (
	"fmt"
	"src/github.com/sotacode/practiceGolangStructure/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando\n", hu.Sexo())
}
