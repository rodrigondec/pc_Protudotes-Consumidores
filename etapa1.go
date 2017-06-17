package main

import (
	"fmt"
)

type Pedido struct {
	id int
	dados string
}

func main() {
	var p Pedido
	p.id = 123
	p.dados = "ai dento"
	fmt.Println(p)
}
