package chamadas

import (
	"fmt"
	"strconv"
)

type Operador struct {
	id         int
	CallCenter *CallCenter
	estaLivre  bool
}

func NovoOperador(id int, callCenter *CallCenter) *Operador {
	return &Operador{
		id:         id,
		CallCenter: callCenter,
		estaLivre:  true,
	}
}

func (operador *Operador) ocupar() {
	fmt.Println("Operador " + strconv.Itoa(operador.id) + " está ocupado")
	operador.estaLivre = false
}

func (operador *Operador) liberar() {
	fmt.Println("Operador " + strconv.Itoa(operador.id) + " está liberado")

	callCenter := operador.CallCenter
	proximoItemDaFila := callCenter.fila.proximoDaFila()
	if proximoItemDaFila != nil {
		callCenter.fila.removeDaFila(proximoItemDaFila)
		*proximoItemDaFila <- operador
	}

	operador.estaLivre = true
}
