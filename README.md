# CallCenter

## Problema:
Crie um programa concorrente em Java ou em Go que ilustre ligações sendo realizadas para um Call Center.

Neste problema teremos um Call Center onde somente 5 operadores estão trabalhando e existem diversos clientes (15 no mínimo) tentando ligar para o mesmo. Cada cliente ao entrar em contato com o Call Center pode ser atendido caso exista um operador disponível ou ele deve entrar numa fila de espera e aguardar até que um atendente esteja disponível. 

Assim que for atendido o cliente poderá ficar em atendimento durante um tempo entre 1 a 5 min (use segundos na sua simulação). Em seguida, ele deve encerrar a ligação e liberar o atendente para o próximo cliente. Faça com que os 15 clientes fiquem indefinidamente tentando entrar em contato com o Call Center. Use mecanismos de sincronização para resolver este problema e lembre-se de imprimir mensagens na saída padrão para indicar o que está acontendo.

Crie uma solução concorrente definindo corretamente as entidades candidatas a serem threads. Utilize a impressão na saída padrão para ilustrar o processo acontecendo e o sleep() para fazer as threads dormirem quando necessário.

## Solução:

A fila do callcenter ser um array (fila) de canais de operador, quando um operador se liberar ele deve escrever ele mesmo no primiro item da fila e remove-lo da fila.

````
type CallCenter struct {
	operadores     *[]*Operador
	fila           *[]*chan *Operador
	mutex          *sync.Mutex
	tempoAleatorio tempoAleatorio
}
````

### Pontos de exclusão mutua

Existem dois pontos onde somente uma routine deve atuar por vez:
- o ponto onde o cliente verifica a disponibilidade de um operador ou espera a liberação de um ocupado
- os pontos de manipulação da fila