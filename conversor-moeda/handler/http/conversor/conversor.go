package conversor

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/adrianamaiaterosendo/conversor-moeda.git/conversor-moeda/model"
	"github.com/adrianamaiaterosendo/conversor-moeda.git/conversor-moeda/service"
)

//ConverteValor - Com os valores digitados, informa a cotação
func ConverteValor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = model.ConversorInput{}
	var body, _ = ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &res)

	cotacao, err := service.BuscarConversao(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(cotacao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
