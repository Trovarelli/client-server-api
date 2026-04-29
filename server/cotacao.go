package main

import "time"

type Cotacao struct {
	ID        int       `json:"-"`
	Code      string    `json:"code"`
	Codein    string    `json:"codein"`
	Name      string    `json:"name"`
	High      string    `json:"high"`
	Low       string    `json:"low"`
	VarBid    string    `json:"varBid"`
	PctChange string    `json:"pctChange"`
	Bid       string    `json:"bid"`
	Ask       string    `json:"ask"`
	Timestamp string    `json:"timestamp"`
	CreatedAt time.Time `json:"createdAt"`
}
