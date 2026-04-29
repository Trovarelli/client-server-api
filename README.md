# Desafio: Client-Server API

Este projeto consiste em um sistema de cotação de dólar composto por um servidor e um cliente em Go.

## Requisitos do Desafio

- O `server.go` deve consumir a API de câmbio (USD-BRL) com timeout de 200ms.
- O `server.go` deve persistir a cotação no SQLite com timeout de 10ms.
- O `server.go` deve expor o endpoint `/cotacao` na porta 8080.
- O `client.go` deve consumir o servidor local com timeout de 300ms.
- O `client.go` deve salvar o valor da cotação (`bid`) em um arquivo `cotacao.txt`.

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
