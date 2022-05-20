package chamadas

import (
	"math/rand"
	"time"
)

type tempoAleatorio interface {
	gerar() time.Duration
}

type TempoAleatorioImplementado struct {
}

func (TempoAleatorioImplementado) gerar() time.Duration {
	minimo := 6
	maximo := 6 * 5
	numeroAleatorio := rand.Intn(maximo-minimo) + minimo
	return time.Duration(numeroAleatorio * 1000000000)
}
