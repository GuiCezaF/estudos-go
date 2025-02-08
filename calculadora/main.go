package main

import (
	"errors"
	"fmt"
)

type Calculadora struct {
	Operador1 float64
	Operador2 float64
}

func (c Calculadora) Somar() float64 {
	return c.Operador1 + c.Operador2
}

func (c Calculadora) Subtrair() float64 {
	return c.Operador1 - c.Operador2
}

func (c Calculadora) Multiplicar() float64 {
	return c.Operador1 * c.Operador2
}
func (c Calculadora) Dividir() (float64, error) {
	if c.Operador2 == 0 {
		return 0, errors.New("Não é possível dividir por zero")
	}
	return c.Operador1 / c.Operador2, nil
}

type Runner struct {
	calculadora *Calculadora
	operacao    string
}

func (r *Runner) SolicitarValores() error {
	fmt.Print("Digite o primeiro valor: ")
	if _, err := fmt.Scan(&r.calculadora.Operador1); err != nil {
		return errors.New("Entrada invalida para o primeiro valor")
	}
	fmt.Print("Digite o segundo valor: ")
	if _, err := fmt.Scan(&r.calculadora.Operador2); err != nil {
		return errors.New("Entrada invalida para o primeiro valor")
	}
	return nil
}

func (r *Runner) SolicitarOperacao() error {
	var operacao string
	fmt.Print("Digite a operação desejada (+, -, *, /): ")
	if _, err := fmt.Scan(&operacao); err != nil {
		return errors.New("Entrada invalida para a operação")
	}

	switch operacao {
	case "+", "-", "*", "/":
		r.operacao = operacao
		return nil
	default:
		return errors.New("Operação inválida")
	}
}

func (r *Runner) ExecutarOperacao() {
	switch r.operacao {
	case "+":
		fmt.Println("Resultado: ", r.calculadora.Somar())
	case "-":
		fmt.Println("Resultado: ", r.calculadora.Subtrair())
	case "*":
		fmt.Println("Resultado: ", r.calculadora.Multiplicar())
	case "/":
		resultado, err := r.calculadora.Dividir()
		if err != nil {
			println(err.Error())
		} else {
			println("Resultado: ", resultado)
		}
	}
}

func (r *Runner) Execute() {
	for {
		if err := r.SolicitarValores(); err != nil {
			println("ERRO: ", err)
			continue
		}
		if err := r.SolicitarOperacao(); err != nil {
			println("ERRO: ", err)
			continue
		}
		r.ExecutarOperacao()
	}
}

func newRunner() *Runner {
	return &Runner{calculadora: &Calculadora{}}
}

func main() {
	fmt.Println("CALCULADORA INFINITA")
	calculadora := &Calculadora{}
	runner := &Runner{calculadora: calculadora}
	runner.Execute()

}
