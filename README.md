# Projeto Ping Pong em Go

Este projeto implementa um simples jogo de "ping pong" entre duas goroutines em Go, utilizando canais para sincronização e comunicação entre elas. A ideia é demonstrar o uso de goroutines e canais de uma forma divertida e educativa.

## Descrição

O programa consiste em duas goroutines que se comunicam através de dois canais, `pings` e `pongs`. Cada goroutine espera por uma mensagem do canal oposto antes de enviar uma resposta de volta pelo seu canal respectivo, alternando as palavras "ping" e "pong" que são exibidas no console.

## Como executar

Para executar este projeto, você precisará ter o [Go instalado](https://golang.org/dl/) em seu sistema. Siga os passos abaixo para baixar e rodar o programa:

1. Clone o repositório para a sua máquina local usando:
  ```bash
  git clone https://github.com/CodeAkio/dio-go-challenge-03.git
  ```

2. Navegue até o diretório do projeto:
  ```bash
  cd dio-go-challenge-03
  ```

3. Execute o programa com o comando:
  ```bash
  go run main.go
  ```
