package main

import (
	"CallCenter/chamadas"
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Iniciando Call Center")
	callCenter := chamadas.NovoCallCenter(chamadas.TempoAleatorioImplementado{})
	for i := 1; i <= 5; i++ {
		fmt.Println("Novo Operador: " + strconv.Itoa(i))
		operador := chamadas.NovoOperador(i, callCenter)
		callCenter.AdicionarOperador(operador)
	}

	for i := 1; i <= 15; i++ {
		cliente := chamadas.NovoCliente(i)
		go cliente.Iniciar(callCenter)
	}

	duration := 2 * time.Minute
	time.Sleep(duration)
}
