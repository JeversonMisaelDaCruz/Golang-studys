package main

import (
	"fmt"
)

type soma struct {
	n1 int
	n2 int
}

func main() {
	var MinhaSoma soma
	fmt.Println("Vamos fazer uma soma insira o primeiro numero")
	fmt.Scanln(&MinhaSoma.n1)
	fmt.Println("Vamos fazer uma soma insira o segundo numero")
	fmt.Scanln(&MinhaSoma.n2)
	fmt.Println("A soma e", MinhaSoma.n2+MinhaSoma.n1)
}
