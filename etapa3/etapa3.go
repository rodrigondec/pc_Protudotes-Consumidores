package main

import (
	"sync"
	"time"
	"fmt"
	"strconv"
	"os"
)

const SEPARADOR = " / " //string para separação de informações
const TAMANHO_BUFFER = 5000 //tamanho do buffer do canal
const LIMITE_PEDIDOS = TAMANHO_BUFFER //limite de pedidos produzidos
const TEMPO_PROCESSAMENTO = 100 //em ms

var is_channel_closed = false //variável de condição se canal está fechado ou não

//estrutura que representa o contador de id_pedido
var contador_id_pedido struct{
	sync.Mutex
	n int
}

var mutex_consumidor_print = sync.Mutex{} //mutex para impressão da retirada do produto do canal

//estrutura que representa um pedido
type Pedido struct {
	id int //identificador
	dados string
}

var consumo sync.WaitGroup //cria grupo de espera de consumo dos pedidos

/* gorotina consumidora que consumirá de um canal
bufferizado com 5000 pedidos */
func consumidor (ch chan Pedido, n int) {
	mutex_consumidor_print.Lock()
	for p := range ch {
		horario_inicio := time.Now()
		fmt.Println("\tConsumidor: " + strconv.Itoa(n) +
			" Retirou o pedido " + strconv.Itoa(p.id) +
			" na hora " + horario_inicio.String())
		mutex_consumidor_print.Unlock()

		time.Sleep(TEMPO_PROCESSAMENTO * time.Millisecond)
		horario_termino := time.Now()

		fmt.Println("\t\tConsumidor: " + strconv.Itoa(n) + SEPARADOR +
			"Pedido: " + strconv.Itoa(p.id) + SEPARADOR +
			"Inicio proc: " + horario_inicio.String() + SEPARADOR +
			"Termino proc: " + horario_termino.String() + SEPARADOR +
			"Duracao: " + horario_termino.Sub(horario_inicio).String())

		mutex_consumidor_print.Lock()
	}
	mutex_consumidor_print.Unlock()
	consumo.Done()
}


/* gorotina produtora que produzirá em um canal
 bufferizado com 5000 pedidos */
func produtor (ch chan Pedido, n int) {
	for {
		var p Pedido
		horario_inicio := time.Now()
		time.Sleep(TEMPO_PROCESSAMENTO * time.Millisecond)

		contador_id_pedido.Lock()
		//condição de parada de produção
		if contador_id_pedido.n > LIMITE_PEDIDOS {
			if !is_channel_closed {
				is_channel_closed = true
				close(ch)
			}
			contador_id_pedido.Unlock()
			return
		}
		//inserção do pedido no canal
		id := contador_id_pedido.n
		contador_id_pedido.n += 1
		p = Pedido{id, "Dados do pedido #" + strconv.Itoa(contador_id_pedido.n)}
		horario_termino := time.Now()
		fmt.Println("Produtor: " + strconv.Itoa(n) + SEPARADOR +
			"Pedido: " + strconv.Itoa(p.id) + SEPARADOR +
			"Inicio proc: " + horario_inicio.String() + SEPARADOR +
			"Termino proc: " + horario_termino.String() + SEPARADOR +
			"Duracao: " + horario_termino.Sub(horario_inicio).String())
		ch <- p

		contador_id_pedido.Unlock()
	}
}

func main() {
	if len(os.Args) == 3 {
		QTD_CONSUMIDORES, _ := strconv.Atoi(os.Args[1])
		QTD_PRODUTORES, _ := strconv.Atoi(os.Args[2])
		contador_id_pedido.n = 1
		ch := make(chan Pedido, TAMANHO_BUFFER) //cria canal

		//executa todos os produtores
		for i := 1; i <= QTD_PRODUTORES; i++ {
			go produtor(ch, i)
		}

		//executa todos os consumidores
		for i := 1; i <= QTD_CONSUMIDORES; i++ {
			consumo.Add(1)
			go consumidor(ch, i)
		}

		//espera termino de execucao de todos os consumidores
		consumo.Wait()
	} else {
		fmt.Println("Numero invalido de argumentos. Requer exatamente 2 parametros enviados:")
		fmt.Println("\n1 - Quantidade de consumidores\n2 - Quantidade de produtores")
	}
}
