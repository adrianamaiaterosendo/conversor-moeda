package service

import (
	"time"

	"github.com/adrianamaiaterosendo/conversor-moeda.git/conversor-moeda/model"
)

func BuscarConversao(valor model.ConversorInput) (model.ConversorRetorno, error) {

	if valor.ValorDolar < 0 {
		return model.ConversorRetorno{}, nil
	}

	valorConvertido := RealizarConversao(valor)

	retorno := model.ConversorRetorno{
		ValorReal:   valorConvertido,
		ValorDolar:  valor.ValorDolar,
		DataCotacao: time.Now(),
	}

	return retorno, nil
}

func RealizarConversao(valor model.ConversorInput) float64 {

	valorConvertido := valor.ValorCompra * valor.ValorDolar

	return valorConvertido
}
