package model

import "time"

type ConversorInput struct {
	ValorDolar float64 `json:"valor_dolar"`
	ValorCompra float64 `json:"valor_compra"`
}

type ConversorRetorno struct {
	ValorReal     float64   `json:"valor_real"`
	ValorDolar    float64   `json:"valor_dolar"`
	DataCotacao time.Time `json:"data_cotacao"`
}
