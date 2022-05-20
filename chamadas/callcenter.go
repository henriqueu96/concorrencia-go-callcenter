package chamadas

import "sync"

type CallCenter struct {
	operadores     *[]*Operador
	fila           Fila
	mutex          *sync.Mutex
	tempoAleatorio tempoAleatorio
}

func NovoCallCenter(tempoAleatorio tempoAleatorio) *CallCenter {
	operadores := &([]*Operador{})
	fila := novaFila()
	return &CallCenter{
		operadores:     operadores,
		fila:           fila,
		tempoAleatorio: tempoAleatorio,
		mutex:          &sync.Mutex{},
	}
}

func (callCenter *CallCenter) AdicionarOperador(operador *Operador) {
	operadores := append(*callCenter.operadores, operador)
	callCenter.operadores = &operadores
}

func (callCenter *CallCenter) novaChamadaDoCliente(cliente *Cliente) {
	callCenter.bloquearNovasChamadas()
	operadorLivre := callCenter.primeiroOperadorLivre()
	if operadorLivre == nil {
		operadorLivre = callCenter.esperaLiberarOperador()
	}
	operadorLivre.ocupar()
	callCenter.liberarNovasChamadas()
	chamada := novaChamada(cliente, operadorLivre, callCenter.tempoAleatorio)
	chamada.iniciar()
}

func (callCenter *CallCenter) primeiroOperadorLivre() *Operador {
	for _, operador := range *callCenter.operadores {
		if operador.estaLivre {
			return operador
		}
	}
	return nil
}

func (callCenter *CallCenter) esperaLiberarOperador() *Operador {
	itemDaFila := make(chan *Operador)
	callCenter.fila.adicionarItemDaFila(&itemDaFila)
	var operador *Operador
	operador = <-itemDaFila
	return operador
}

func (callCenter *CallCenter) bloquearNovasChamadas() {
	callCenter.mutex.Lock()
}

func (callCenter *CallCenter) liberarNovasChamadas() {
	callCenter.mutex.Unlock()
}
