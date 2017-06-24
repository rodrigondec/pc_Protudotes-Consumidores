# Etapa 3

## Descrição

Nessa etapa há produtores, i.e. clientes fazendo pedidos. O servidor possui um buffer com espaço de armazenamento para 5000 pedidos. Cria-se gorotinas produtoras e consumidoras \(a quantidade foi especificada via linha de comando; vide [instruções de execução](1-instrucoes.md)\).

As gorotinas produtoras serão responsáveis por criar pedidos e inserí-los no buffer seguindo a ordem de criação para a inserção no buffer. Para simular o processamento, cada uma dessas gorotinas irão dormir por um tempo X antes de adicionar o pedido no buffer. A cada execução do programa serão criados, no máximo, 5000 pedidos.

Já as gorotinas consumidoras irão retirar os pedidos do buffer, na ordem de criação, e processá-los até que o buffer seja esvaziado. Após retirar um pedido do buffer, a gorotina em questão irão por um tempo X. Esse tempo é necessário para simular processamento, e.g. alterar o banco de dados da empresa.

Note que, diferentemente da etapa 2, os pedidos inseridos no buffer seguem a ordem de criação e a retirada dos pedidos do canal para o consumo também obedece essa ordem.

Foram realizados benchmarkings com tempo X de 100ms.

## Benchmarking

A seguinte tabela mostra uma análise estatística do tempo necessário para executar o algoritmo para diferentes quantidades de gorotinas simultaneamente \(representadas na primeira coluna\). O código foi rodado somente nos casos em que a quantidade de consumidores iguala a quantidade de produtores, ou seja: \[\(1, 1\), \(5, 5\), \(10, 10\), \(50, 50\), \(100, 100\), \(500, 500\), \(1000, 1000\), \(5000, 5000\)\]. O algoritmo foi executado 80 vezes: 10 vezes para cada entrada \(linha da tabela\), conforme requisitado na [especificação do projeto](../Trabalho-Go.pdf).

### 100ms

| Qt consumidores | Qt produtores | Tempo Mínimo | Tempo Médio | Tempo Máximo | Desvio Padrão |
| :---: | :---: | :---: | :---: | :---: | :---: |
| 1 | 1 | 502.066 | 502.2001 | 502.434 | 0.1082019 |
| 5 | 5 | 101.038 | 101.0836 | 101.118 | 0.02622637 |
| 10 | 10 | 50.9 | 50.9147 | 50.932 | 0.010133 |
| 50 | 50 | 10.825 | 10.8536 | 10.909 | 0.02912883 |
| 100 | 100 | 6.301 | 6.5717 | 6.831 | 0.1994493 |
| 500 | 500 | 5.791 | 6.205 | 6.531 | 0.2706851 |
| 1000 | 1000 | 6.179 | 6.398 | 6.482 | 0.09271462 |
| 5000 | 5000 | 6.4 | 6.5299 | 6.661 | 0.1016863 |





