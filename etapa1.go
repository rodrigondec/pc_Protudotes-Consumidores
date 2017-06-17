package main

import (
	"fmt"
	"strconv"
	"time"
)

const BUFFER_SIZE = 10

//estrutura que representa um pedido
type Pedido struct {
	id int //identificador
	dados string
}

/*gorotina consumidora que consumira um canal
bufferizado com 5000 pedidos*/
func consumidor (ch chan Pedido) {
	for p := range ch {
		fmt.Println(p)
	}
}

func main() {
	var p Pedido
	ch := make(chan Pedido, BUFFER_SIZE) //cria canal

	for i := 1; i <= BUFFER_SIZE; i++ {
		p = Pedido{i, "Dados do pedido #" + strconv.Itoa(i)}
		ch <- p
	}
	close(ch)
	go consumidor(ch)
	time.Sleep(1)
	
}
