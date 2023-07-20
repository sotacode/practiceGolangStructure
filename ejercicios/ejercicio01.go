package ejercicios

import "strconv"

func ConviertoAEntero(parametro string) (int, string) {
	numero, err := strconv.Atoi(parametro)
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}
	if numero > 100 {
		return numero, "Es mayor que 100"
	} else {
		return numero, "Es menor o igual que 100"
	}
}
