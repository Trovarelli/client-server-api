package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalf("Erro ao criar requisição: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("Erro: Timeout de 300ms atingido na requisição ao servidor")
		} else {
			log.Printf("Erro ao realizar requisição: %v", err)
		}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Servidor retornou erro: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler resposta: %v", err)
	}

	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}

	if cotacao.Bid == "" {
		log.Fatal("Campo 'bid' não encontrado na resposta")
	}

	err = saveToFile(cotacao.Bid)
	if err != nil {
		log.Fatalf("Erro ao salvar arquivo: %v", err)
	}

	fmt.Printf("Cotação salva com sucesso: Dólar: %s\n", cotacao.Bid)
}

func saveToFile(bid string) error {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	content := fmt.Sprintf("Dólar: %s", bid)
	_, err = file.WriteString(content)
	return err
}
