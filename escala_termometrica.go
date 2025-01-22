package main

import "fmt"

func main() {

	const ebulicao_K = 373

	celsius := ebulicao_K - 273
	fmt.Printf("O ponto de ebulição da água é %f K, que corresponde a %f °C.", ebulicao_K, celsius)

}
