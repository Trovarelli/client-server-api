# Desafio: Client-Server API

Este projeto consiste em um sistema de cotação de dólar composto por um servidor e um cliente em Go.

## Como Rodar

### 1. Iniciar o Servidor
Navegue até a pasta `server` e execute:
```bash
go run .
```
O servidor estará rodando em `http://localhost:8080`.

### 2. Executar o Cliente
Em outro terminal, navegue até a pasta `client` e execute:
```bash
go run .
```
O valor da cotação será salvo no arquivo `client/cotacao.txt`.

## Tecnologias Utilizadas
- Go (Golang)
- SQLite (modernc.org/sqlite - pure Go implementation)
- Context para controle de timeouts.
