package main

import "fmt"

func main() {

	const ebulicao_K = 373

	celsius := ebulicao_K - 273
	fmt.Printf("O ponto de ebulição da água é %d K, que corresponde a %d °C.", ebulicao_K, celsius)

}
