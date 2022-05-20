package chamadas

import "sync"

type Fila struct {
	mutex    *sync.Mutex
	channels *[]*chan *Operador
}

func novaFila() Fila {
	channels := &([]*chan *Operador{})
	return Fila{
		channels: channels,
		mutex:    &sync.Mutex{},
	}
}

func (fila *Fila) adicionarItemDaFila(itemDaFila *chan *Operador) {
	fila.mutex.Lock()
	novaFila := append(*fila.channels, itemDaFila)
	fila.channels = &novaFila
	fila.mutex.Unlock()
}

func (fila *Fila) removeDaFila(itemASerRemovido *chan *Operador) {
	fila.mutex.Lock()
	var novaFila []*chan *Operador
	for _, itemDaFila := range *fila.channels {
		if itemDaFila != itemASerRemovido {
			novaFila = append(novaFila, itemDaFila)
		}
	}
	fila.channels = &novaFila
	fila.mutex.Unlock()
}

func (fila *Fila) proximoDaFila() *chan *Operador {
	fila.mutex.Lock()
	channels := *fila.channels
	if len(channels) == 0 {
		return nil
	}
	itemDaFila := channels[0]
	fila.mutex.Unlock()
	return itemDaFila
}
