package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

type ExchangeRateResponse struct {
	USDBRL Cotacao `json:"USDBRL"`
}

func Router(db *sql.DB) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cotacao", GetCotacaoHandler(db))

	log.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", mux)
}

func GetCotacaoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxAPI, cancelAPI := context.WithTimeout(r.Context(), 200*time.Millisecond)
		defer cancelAPI()

		cotacao, err := FetchCotacao(ctxAPI)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				log.Println("Erro: Timeout de 200ms atingido na API")
				http.Error(w, "API timeout", http.StatusRequestTimeout)
			} else {
				log.Printf("Erro ao buscar cotação: %v", err)
				http.Error(w, "Erro ao buscar cotação", http.StatusInternalServerError)
			}
			return
		}

		ctxDB, cancelDB := context.WithTimeout(r.Context(), 10*time.Millisecond)
		defer cancelDB()

		err = InsertCotacao(ctxDB, db, cotacao)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				log.Println("Erro: Timeout de 10ms atingido no banco de dados")
			} else {
				log.Printf("Erro ao salvar no banco: %v", err)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cotacao)
	}
}

func FetchCotacao(ctx context.Context) (*Cotacao, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res ExchangeRateResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	c := res.USDBRL
	c.CreatedAt = time.Now()
	return &c, nil
}
