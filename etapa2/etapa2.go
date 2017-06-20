package main

import (
	"sync"
	"time"
	"fmt"
	"strconv"
	"os"
)

const SEPARADOR = " / "
const TAMANHO_BUFFER = 5000
var id_pedido = 1
var pedidos_terminados = false
var is_channel_closed = false

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
		horario_inicio := time.Now()
		time.Sleep(500 * time.Millisecond)
		horario_termino := time.Now()

		fmt.Println("Consumidor: " + strconv.Itoa(n) + SEPARADOR +
			"Pedido: " + strconv.Itoa(p.id) + SEPARADOR +
			"Inicio proc: " + horario_inicio.String() + SEPARADOR +
			"Termino proc: " + horario_termino.String() + SEPARADOR +
			"Duracao: " + horario_termino.Sub(horario_inicio).String())
	}
	wg.Done()
}


/* gorotina produtora que produzirá em um canal
 bufferizado com 5000 pedidos */
func produtor (ch chan Pedido, n int, mutex_qt *sync.Mutex, mutex_id *sync.Mutex) {
	/*
    defer func() {
        if rec := recover(); rec != nil  {
        	fmt.Println("Recover at produtor:", rec)
        }
    }()*/

	for {
		var p Pedido
		horario_inicio := time.Now()
		time.Sleep(500 * time.Millisecond)
		
		if id_pedido > TAMANHO_BUFFER {
			if !is_channel_closed {
				is_channel_closed = true
				close(ch)
			}
			return
		}
		p = Pedido{id_pedido, "Dados do pedido #" + strconv.Itoa(id_pedido)}
		id_pedido += 1
		ch <- p
		horario_termino := time.Now()


		fmt.Println("\tProdutor: " + strconv.Itoa(n) + SEPARADOR +
			"Pedido: " + strconv.Itoa(p.id) + SEPARADOR +
			"Inicio proc: " + horario_inicio.String() + SEPARADOR +
			"Termino proc: " + horario_termino.String() + SEPARADOR +
			"Duracao: " + horario_termino.Sub(horario_inicio).String())

	}
}


func main() {
	if len(os.Args) == 3 {
		QTD_CONSUMIDORES, _ := strconv.Atoi(os.Args[1])
		QTD_PRODUTORES, _ := strconv.Atoi(os.Args[2])
		ch := make(chan Pedido, TAMANHO_BUFFER) //cria canal
		var m_qt_atividades = &sync.Mutex{}
		var m_id_atividades = &sync.Mutex{}

		//executa todos os produtores
		for i := 1; i <= QTD_PRODUTORES; i++ {
			go produtor(ch, i, m_qt_atividades, m_id_atividades)
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
		fmt.Println("1 - Tamanho do buffer / quantidade de pedidos\n2 - Quantidade de gorotinas")
	}
}
