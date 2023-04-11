package main

import "fmt"

func main() {
	x := 77                             // declaración de variables
	fmt.Println("Descubre el mundo...") // salida de texto
	nb, err := fmt.Println("Cambios locales ", x)
	// ...interface() cantidad variable de parametros
	// _,_ := hace referencia a la variable que no se utilizara
	if err != nil {
		fmt.Println("...", nb, err)
	}
	//fmt.Println("...", nb, err)
}

// ejecutar codigo $ go run ejercicio.go
