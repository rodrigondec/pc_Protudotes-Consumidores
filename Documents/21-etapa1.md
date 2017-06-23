# Etapa 1

## Descrição

Nessa etapa não há produtores, i.e. não há clientes fazendo pedidos. O servidor possui um buffer com espaço de armazenamento para 5000 pedidos. Essa quantidade de pedidos é criada previamente e o buffer é preenchido.

Em seguida, são chamadas n gorotinas para processá-los até que o buffer seja esvaziado. A quantidade n de gorotinas é enviada por parâmetro via linha de comando \(vide [instruções de execução](1-instrucoes.md)\). Após retirar um pedido do buffer, a gorotina em questão irá dormir um tempo X. Esse tempo é necessário para simular processamento, e.g. alterar o banco de dados da empresa.

Foram realizados benchmarkings com tempo X de 500ms e 100ms.

## Benchmarking

A seguinte tabela mostra uma análise estatística do tempo necessário para executar o algoritmo para diferentes quantidades de gorotinas simultaneamente \(representadas na primeira coluna\). O algoritmo foi executado 80 vezes: 10 vezes para cada entrada \(linha da tabela\), conforme requisitado na [especificação do projeto](../Trabalho-Go.pdf).

### 100ms

| Qt consumidores | Tempo Mínimo | Tempo Médio | Tempo Máximo | Desvio Padrão |
| :---: | :---: | :---: | :---: | :---: |
| 1 |  |  |  |  |
| 5 |  |  |  |  |
| 10 |  |  |  |  |
| 50 |  |  |  |  |
| 100 |  |  |  |  |
| 500 |  |  |  |  |
| 1000 |  |  |  |  |
| 5000 |  |  |  |  |

### 500ms

| Qt consumidores | Tempo Mínimo | Tempo Médio | Tempo Máximo | Desvio Padrão |
| :---: | :---: | :---: | :---: | :---: |
| 1 | 2502 | 2503 | 2504 | 0.4714045 |
| 5 | 501 | 501.1 | 502 | 0.3162278 |
| 10 | 250 | 250.7 | 251 | 0.4830459 |
| 50 | 50 | 50.8 | 51 | 0.421637 |
| 100 | 25 | 25.7 | 26 | 0.4830459 |
| 500 | 5 | 5.8 | 6 | 0.421637 |
| 1000 | 3 | 3.6 | 4 | 0.5163978 |
| 5000 | 2 | 2.8 | 3 | 0.421637 |



