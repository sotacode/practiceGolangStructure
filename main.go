package main

import (
	"src/github.com/sotacode/practiceGolangStructure/webserver"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	/* estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows")
	} else {
		fmt.Println("Esto es windows: ", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("darwin")
	default:
		fmt.Printf("%s \n", os)
	} */
	/* numero, mensaje := ejercicios.ConviertoAEntero("34ds")
	fmt.Println(numero)
	fmt.Println(mensaje) */
	//teclado.IngresoNumeros()
	//iteraciones.Iterar()
	/* texto := ejercicios.TablaMultiplicar()
	fmt.Println(texto) */
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.Exponencia(2)
	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()
	/* Pedro := new(modelos.Hombre)
	ejer_interfaces.HumanosRespirando(Pedro)
	Maria := new(modelos.Mujer)
	ejer_interfaces.HumanosRespirando(Maria) */
	//defer_example.EjemploPanic()
	/* canal1 := make(chan bool)
	go goroutines.MiNombreLentooo("Sota", canal1)
	fmt.Println("Estoy aqui")
	<-canal1 //esto seria una especie de await de la go routine */
	webserver.MiWebServer()

}
