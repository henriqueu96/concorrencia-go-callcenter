package chamadas

import (
	"fmt"
	"strconv"
	"time"
)

type Chamada struct {
	cliente  *Cliente
	operador *Operador
	tempo    time.Duration
}

func novaChamada(cliente *Cliente, operador *Operador, tempoAleatorio tempoAleatorio) Chamada {
	return Chamada{
		cliente:  cliente,
		operador: operador,
		tempo:    tempoAleatorio.gerar(),
	}
}

func (chamada *Chamada) iniciar() {
	fmt.Println("Cliente " + strconv.Itoa(chamada.cliente.id) + " sendo atendido")
	time.Sleep(chamada.tempo)
	chamada.operador.liberar()
}
