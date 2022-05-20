package chamadas

type Cliente struct {
	id int
}

func NovoCliente(id int) Cliente {
	return Cliente{
		id: id,
	}
}

func (cliente Cliente) Iniciar(callCenter *CallCenter) {
	for {
		cliente.ligarParaCallCenter(callCenter)
	}
}

func (cliente *Cliente) ligarParaCallCenter(callCenter *CallCenter) {
	callCenter.novaChamadaDoCliente(cliente)
}
