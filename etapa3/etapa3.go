package main

import (
	"sync"
	"time"
	"fmt"
	"strconv"
	"os"
)

const SEPARADOR = " / "
const TAMANHO_BUFFER = 10

var is_channel_closed = false

var id_pedido struct{
	sync.Mutex
	n int
}

var pedido_processado [TAMANHO_BUFFER]sync.WaitGroup

var m_print_p = sync.Mutex{}
var m_print_c = sync.Mutex{}

//estrutura que representa um pedido
type Pedido struct {
	id int //identificador
	dados string
}

var wg sync.WaitGroup //cria grupo de espera

/* gorotina consumidora que consumirá de um canal
bufferizado com 5000 pedidos */
func consumidor (ch chan Pedido, n int) {
	for p := range ch {
		index_pedido := p.id-1
		if index_pedido != 0{
			fmt.Println("Consumidor: " + strconv.Itoa(n) + SEPARADOR +
				"Vai esperar o pedido " + strconv.Itoa((p.id-1)) + " ser processado")
			pedido_processado[(index_pedido-1)].Wait()
		}

		horario_inicio := time.Now()
		time.Sleep(500 * time.Millisecond)
		horario_termino := time.Now()

		fmt.Println("Consumidor: " + strconv.Itoa(n) + SEPARADOR +
			"Pedido: " + strconv.Itoa(p.id) + SEPARADOR +
			"Inicio proc: " + horario_inicio.String() + SEPARADOR +
			"Termino proc: " + horario_termino.String() + SEPARADOR +
			"Duracao: " + horario_termino.Sub(horario_inicio).String())

		pedido_processado[(index_pedido)].Done()
	}
	wg.Done()
}


/* gorotina produtora que produzirá em um canal
 bufferizado com 5000 pedidos */
func produtor (ch chan Pedido, n int) {
	for {
		var p Pedido
		horario_inicio := time.Now()
		time.Sleep(500 * time.Millisecond)
		if id_pedido.n > TAMANHO_BUFFER {
			if !is_channel_closed {
				is_channel_closed = true
				close(ch)
			}
			return
		}

		id_pedido.Lock()
		id := id_pedido.n
		id_pedido.n += 1
		p = Pedido{id, "Dados do pedido #" + strconv.Itoa(id_pedido.n)}
		ch <- p
		horario_termino := time.Now()

		fmt.Println("\tProdutor: " + strconv.Itoa(n) + SEPARADOR +
			"Pedido: " + strconv.Itoa(p.id) + SEPARADOR +
			"Inicio proc: " + horario_inicio.String() + SEPARADOR +
			"Termino proc: " + horario_termino.String() + SEPARADOR +
			"Duracao: " + horario_termino.Sub(horario_inicio).String())
		id_pedido.Unlock()
	}
}


func main() {
	if len(os.Args) == 3 {
		QTD_CONSUMIDORES, _ := strconv.Atoi(os.Args[1])
		QTD_PRODUTORES, _ := strconv.Atoi(os.Args[2])
		id_pedido.n = 1
		ch := make(chan Pedido, TAMANHO_BUFFER) //cria canal

		for i := 0; i < TAMANHO_BUFFER; i++{
			pedido_processado[i].Add(1)
		}

		//executa todos os produtores
		for i := 1; i <= QTD_PRODUTORES; i++ {
			go produtor(ch, i)
		}

		//executa todos os consumidores
		for i := 1; i <= QTD_CONSUMIDORES; i++ {
			wg.Add(1)
			go consumidor(ch, i)
		}

		//espera termino de execucao de todos os consumidores
		wg.Wait()
	} else {
		fmt.Println("Numero invalido de argumentos. Requer exatamente 2 parametros enviados:\n")
		fmt.Println("1 - Quantidade de consumidores\n2 - Quantidade de produtores")
	}
}
