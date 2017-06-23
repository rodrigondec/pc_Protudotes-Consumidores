# Etapa 2

## Descrição

Nessa etapa há produtores, i.e. clientes fazendo pedidos. O servidor possui um buffer com espaço de armazenamento para 5000 pedidos. Cria-se gorotinas produtoras e consumidoras \(a quantidade foi especificada via linha de comando; vide [instruções de execução](1-instrucoes.md)\).

As gorotinas produtoras serão responsáveis por criar pedidos e inserí-los no buffer. Para simular o processamento, cada uma dessas gorotinas irão dormir por um tempo X antes de adicionar o pedido no buffer. A cada execução do programa serão criados, no máximo, 5000 pedidos.

Já as gorotinas consumidoras irão retirar os pedidos do buffer e processá-los até que o buffer seja esvaziado. Após retirar um pedido do buffer, a gorotina em questão irão por um tempo X. Esse tempo é necessário para simular processamento, e.g. alterar o banco de dados da empresa.

Note que, diferentemente da etapa 1, os pedidos são gerados on-the-go e, não necessariamente, o buffer irá ficar cheio.

Para compararmos a etapa 2 com a etapa 1, o tempo X de processamento tanto dos produtores quanto consumidores será de meio segundo \(500ms\). Enquanto para compararmos a etapa 2 com a etapa 3, o tempo X de processamento tanto de produtores quanto de consumidores será de 0,05 segundo \(50ms\).

## Benchmarking

A seguinte tabela mostra uma análise estatística do tempo necessário para executar o algoritmo para diferentes quantidades de gorotinas simultaneamente \(representadas na primeira coluna\). O código foi rodado somente nos casos em que a quantidade de consumidores iguala a quantidade de produtores, ou seja: \[\(1, 1\), \(5, 5\), \(10, 10\), \(50, 50\), \(100, 100\), \(500, 500\), \(1000, 1000\), \(5000, 5000\)\]. O algoritmo foi executado 80 vezes: 10 vezes para cada entrada \(linha da tabela\), conforme requisitado na [especificação do projeto](../Trabalho-Go.pdf).

### Tabela 500ms

| Qt consumidores | Qt produtores | Tempo Mínimo | Tempo Médio | Tempo Máximo | Desvio Padrão |
| :---: | :---: | :---: | :---: | :---: | :---: |
| 1 | 2502 | 2502.6 | 2503 | 0.5163978 |
| 5 | 501 | 501.5 | 502 | 0.5270463 |
| 10 | 251 | 251.3 | 252 | 0.4830459 |
| 50 | 51 | 51.3 | 52 | 0.4830459 |
| 100 | 26 | 26.2 | 27 | 0.4830459 |
| 500 | 6 | 6.9 | 7 | 0.3162278 |
| 1000 | 6 | 6.5 | 7 | 0.5270463 |
| 5000 | 6 | 6.6 | 7 | 0.5163978 |

### Tabela 50ms

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



