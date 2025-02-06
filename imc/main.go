package main

import "fmt"

func calcIMC(peso, altura float64) (float64, error) {
	if peso <= 0 || altura <= 0 {
		return 0, fmt.Errorf("Peso e altura devem ser maiores que zero\n")
	}
	imc := peso / (altura * altura)
	return imc, nil
}

func main() {
	var peso, altura float64

	fmt.Println("Calculo de IMC")

	fmt.Print("Digite o peso (em Kg): ")
	fmt.Scan(&peso)

	fmt.Print("Digite a altura (em Metros): ")
	fmt.Scan(&altura)

	imc, err := calcIMC(peso, altura)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Seu IMC: %.2f \n", imc)
	}
}
