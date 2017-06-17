package main

import (
	"sync"
	"time"
	"fmt"
	"strconv"
	"os"
)

const SEPARADOR = " / "
var id_pedido = 1
var pedidos_terminados = false

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
func produtor (ch chan Pedido, n int, mutex_qt *sync.Mutex, mutex_id *sync.Mutex){
	for {
		mutex_qt.Lock()
		if pedidos_terminados{
			mutex_qt.Unlock()
			break
		}
		if id_pedido == 4999{
			pedidos_terminados = true
		}
		mutex_qt.Unlock()

		var p Pedido
		horario_inicio := time.Now()
		time.Sleep(500 * time.Millisecond)
		mutex_id.Lock()
		p = Pedido{id_pedido, "Dados do pedido #" + strconv.Itoa(id_pedido)}
		id_pedido += 1
		mutex_id.Unlock()
		ch <- p
		horario_termino := time.Now()



		fmt.Println("Produtor: " + strconv.Itoa(n) + SEPARADOR +
			"Pedido: " + strconv.Itoa(p.id) + SEPARADOR +
			"Inicio proc: " + horario_inicio.String() + SEPARADOR +
			"Termino proc: " + horario_termino.String() + SEPARADOR +
			"Duracao: " + horario_termino.Sub(horario_inicio).String())

		if pedidos_terminados{
			close(ch)
		}


	}
}


func main() {
	if len(os.Args) == 3 {
		var TAMANHO_BUFFER = 5000
		QTD_CONSUMIDORES, _ := strconv.Atoi(os.Args[2])
		QTD_PRODUTORES, _ := strconv.Atoi(os.Args[3])
		ch := make(chan Pedido, TAMANHO_BUFFER) //cria canal
		var m_qt_atividades = &sync.Mutex{}
		var m_id_atividades = &sync.Mutex{}

		//executa todos os produtores
		for i := 1; i <= QTD_PRODUTORES; i++ {
			//wg.Add(1)
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
