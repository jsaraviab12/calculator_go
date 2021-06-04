package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jsaraviab12/calculadora/operators"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	resultado := 0
	if strings.Contains(operacion, "+") {
		resultado = operators.Plus(operacion)
	} else if strings.Contains(operacion, "-") {
		resultado = operators.Sub(operacion)
	} else if strings.Contains(operacion, "*") {
		resultado = operators.Multi(operacion)
	} else if strings.Contains(operacion, "/") {
		resultado = operators.Div(operacion)
	}

	fmt.Println(resultado)

}
