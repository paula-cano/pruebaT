package main

import (
	"fmt"
	"Backend/request"
)

func main() {
	// Llamar a la función Request para capturar los valores de retorno
	resp, err := request.Request()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Respuesta recibida:")
	fmt.Println(resp)
}
