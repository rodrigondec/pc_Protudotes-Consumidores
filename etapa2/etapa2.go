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
func produtor (ch chan Pedido, n int){
	for {
		//value, open := <- ch
		//if !open{
		//
		//}

		var p Pedido
		horario_inicio := time.Now()
		p = Pedido{id_pedido, "Dados do pedido #" + strconv.Itoa(id_pedido)}
		id_pedido += 1
		ch <- p
		horario_termino := time.Now()

		fmt.Println("Produtor: " + strconv.Itoa(n) + SEPARADOR +
			"Pedido: " + strconv.Itoa(p.id) + SEPARADOR +
			"Inicio proc: " + horario_inicio.String() + SEPARADOR +
			"Termino proc: " + horario_termino.String() + SEPARADOR +
			"Duracao: " + horario_termino.Sub(horario_inicio).String())

		time.Sleep(500 * time.Millisecond)
	}

// Precisa ter um wait group para os produtores? acho que não pq produzem indefinidamente né?
}


func main() {
	if len(os.Args) == 4 {
		TAMANHO_BUFFER, _ := strconv.Atoi(os.Args[1])
		QTD_CONSUMIDORES, _ := strconv.Atoi(os.Args[2])
		//QTD_PRODUTORES, _ := strconv.Atoi(os.Args[3])
		ch := make(chan Pedido, TAMANHO_BUFFER) //cria canal

		//executa todos os produtores
		for i := 1; i <= QTD_PRODUTORES; i++ {
			//wg.Add(1)
			go produtor(ch, i)
		}

		//close(ch) //fecha o canal
		//ch <- Pedido{1, "Dados do pedido #" + strconv.Itoa(1)}

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
