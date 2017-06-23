# O problema

O problema, resumidamente, consiste em simular o servidor de uma empresa de compras, por exemplo. Essa empresa recebe vários pedidos simultaneamente e o servidor deve ser capaz de processá-los. O objetivo principal do trabalho é compararmos a eficiência de uma solução sequencial (servidor executando uma única gorotina) com diversas soluções concorrentes (servidor executando quantidades diferentes de gorotinas).

Cada etapa adiciona, gradativamente, funcionalidades que estarão presentes na versão final. Ao final de cada etapa é realizado um benchmarking para análise de eficiência de execução. Todos os testes foram rodados em um Raspberry Pi 3 Modelo B.

Para mais detalhes, vide a [especificação do projeto](../Trabalho-Go.pdf).
