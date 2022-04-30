package http

import (
	"log"
	"net/http"

	"github.com/adrianamaiaterosendo/conversor-moeda.git/conversor-moeda/handler/http/conversor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Init é onde eu tenho as chamadas http
func Init() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/transactions", conversor.ConverteValor)

	//cria metricas da aplicação onde podemos customizar
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
