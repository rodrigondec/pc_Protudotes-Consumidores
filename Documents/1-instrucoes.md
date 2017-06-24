# Instruções

## Requerimentos

* [Compilador de GO](https://golang.org/doc/install)

## Executando

Para executar o programa de cada uma das etapas basta executar o comando `go run` seguido pelo nome do arquivo. Note, porém, que cada etapa receberá parâmetros distintos.

### Etapa 1

Supondo que o usuário esteja na mesma pasta do arquivo \(prog-conc\_Protudotes-Consumidores/etapa1/\), o seguinte comando deve ser executado.

```
go run etapa1.go <qtd_consumidores>

<qtd_consumidores> Quantidade de gorotinas que irão consumir concorrentemente o buffer já cheio.
```

Para mais detalhes sobre o funcionamento desse algoritmo acesse a [documentação da etapa1](21-etapa1.md)

### Etapa 2

Supondo que o usuário esteja no diretório pai do arquivo \(prog-conc\_Protudotes-Consumidores/\), o seguinte comando deve ser executado.

```
go run etapa2/etapa2.go <qtd_consumidores> <qtd_produtores>

<qtd_consumidores> Quantidade de gorotinas consumidoras. Irão retirar pedidos da fila e processá-los
<qtd_produtores> Quantidade de gorotinas produtoras. Irão inserir pedidos na fila.
```

Para mais detalhes sobre o funcionamento desse algoritmo acesse a [documentação da etapa2](22-etapa2.md)

### Etapa 3

Supondo que o usuário esteja no diretório pai do arquivo \(prog-conc\_Protudotes-Consumidores/\), o seguinte comando deve ser executado.

```
go run etapa3/etapa3.go <qtd_consumidores> <qtd_produtores>

<qtd_consumidores> Quantidade de gorotinas consumidoras. Irão retirar pedidos da fila e processá-los
<qtd_produtores> Quantidade de gorotinas produtoras. Irão inserir pedidos na fila.
```

Para mais detalhes sobre o funcionamento desse algoritmo acesse a [documentação da etapa3](/Documents/23-etapa3.md)

