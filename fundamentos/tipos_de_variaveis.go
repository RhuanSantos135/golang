package main

import (
	"fmt"
)

func main() {

	nomes := []string{}

	var scan string

	fmt.Println("Coloque o seu nome")
	fmt.Scanf("%s", &scan)

	nomes = append(nomes, scan)
	fmt.Println(nomes)

}
