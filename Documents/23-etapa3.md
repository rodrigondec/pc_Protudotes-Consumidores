# Etapa 3

## Descrição

Nessa etapa há produtores, i.e. clientes fazendo pedidos. O servidor possui um buffer com espaço de armazenamento para 5000 pedidos. Cria-se gorotinas produtoras e consumidoras \(a quantidade foi especificada via linha de comando; vide [instruções de execução](1-instrucoes.md)\).

As gorotinas produtoras serão responsáveis por criar pedidos e inserí-los no buffer seguindo uma ordem de criação para a inserção no buffer. Para simular o processamento, cada uma dessas gorotinas irão dormir por um tempo X antes de adicionar o pedido no buffer. A cada execução do programa serão criados, no máximo, 5000 pedidos.

Já as gorotinas consumidoras irão retirar os pedidos do buffer e processá-los, na ordem de criação, até que o buffer seja esvaziado. Após retirar um pedido do buffer, a gorotina em questão irão por um tempo X. Esse tempo é necessário para simular processamento, e.g. alterar o banco de dados da empresa.

Note que, diferentemente da etapa 2, os pedidos inseridos no buffer seguem a ordem de criação e a retirada dos pedidos do canal para o consumo também obedece essa ordem.

Foram realizados benchmarkings com tempo X de 100ms.

## Benchmarking

A seguinte tabela mostra uma análise estatística do tempo necessário para executar o algoritmo para diferentes quantidades de gorotinas simultaneamente \(representadas na primeira coluna\). O código foi rodado somente nos casos em que a quantidade de consumidores iguala a quantidade de produtores, ou seja: \[\(1, 1\), \(5, 5\), \(10, 10\), \(50, 50\), \(100, 100\), \(500, 500\), \(1000, 1000\), \(5000, 5000\)\]. O algoritmo foi executado 80 vezes: 10 vezes para cada entrada \(linha da tabela\), conforme requisitado na [especificação do projeto](../Trabalho-Go.pdf).

| Qt consumidores | Qt produtores | Tempo Mínimo | Tempo Médio | Tempo Máximo | Desvio Padrão |
| :---: | :---: | :---: | :---: | :---: | :---: |
| 1 | 1 |  |  |  |  |
| 5 | 5 |  |  |  |  |
| 10 | 10 |  |  |  |  |
| 50 | 50 |  |  |  |  |
| 100 | 100 |  |  |  |  |
| 500 | 500 |  |  |  |  |
| 1000 | 1000 |  |  |  |  |
| 5000 | 5000 |  |  |  |  |



