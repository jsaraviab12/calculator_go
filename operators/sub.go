package operators

import (
	"fmt"
	"strconv"
	"strings"
)

//Sub suma
func Sub(operacion string) int {
	valores := strings.Split(operacion, "-")
	resultado := 0

	for i := 0; i < len(valores); i++ {
		num, error := strconv.Atoi(valores[i])

		if error != nil {
			fmt.Println("Tiene que ingresar un numero entero o solo realizar con un operador")
		} else {
			if resultado == 0 {
				resultado = num
			} else {
				resultado -= num
			}

		}

	}
	return resultado
}
