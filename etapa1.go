package main

import (
	"fmt"
	"strconv"
	"time"
	"sync"
)

const TAMANHO_BUFFER = 500
const QTD_CONSUMIDORES = 100
const SEPARADOR = " / "

//estrutura que representa um pedido
type Pedido struct {
	id int //identificador
	dados string
}

var wg sync.WaitGroup //cria grupo de espera

/*gorotina consumidora que consumira um canal
bufferizado com 5000 pedidos*/
func consumidor (ch chan Pedido, n int) {
	for p := range ch {
		horario_inicio := time.Now()
		time.Sleep(500 * time.Millisecond)
		horario_termino := time.Now()

		fmt.Println("Agente: " + strconv.Itoa(n) + SEPARADOR +
			"Pedido: " + strconv.Itoa(p.id) + SEPARADOR +
			"Inicio proc: " + horario_inicio.String() + SEPARADOR +
			"Termino proc: " + horario_termino.String() + SEPARADOR +
			"Duracao: " + horario_termino.Sub(horario_inicio).String())
	}
	wg.Done()
}

func main() {
	var p Pedido
	ch := make(chan Pedido, TAMANHO_BUFFER) //cria canal

	//loop adiciona pedidos no canal
	for i := 1; i <= TAMANHO_BUFFER; i++ {
		p = Pedido{i, "Dados do pedido #" + strconv.Itoa(i)}
		ch <- p
	}
	close(ch) //fecha o canal

	//executa todos os consumidores
	for i := 0; i < QTD_CONSUMIDORES; i++ {
		wg.Add(1)
		go consumidor(ch, i)
	}

	//espera termino de execucao de todos os consumidores
	wg.Wait()

}