package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Endereco representa os dados retornados pelo ViaCEP.
type Endereco struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

// comentário 1 corrigido
// cepHandler trata as requisições para consulta de CEP.
func cepHandler(w http.ResponseWriter, r *http.Request) {
	// Extrai o CEP a partir dos parâmetros da URL.
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "O parâmetro 'cep' é obrigatório", http.StatusBadRequest)
		return
	}

	// Monta a URL da API do ViaCEP e faz a requisição.
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Erro ao acessar o ViaCEP", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erro ao ler a resposta do ViaCEP", http.StatusInternalServerError)
		return
	}

	// Converte o JSON para a estrutura Endereco.
	var endereco Endereco
	if err := json.Unmarshal(body, &endereco); err != nil {
		http.Error(w, "Erro ao processar o JSON", http.StatusInternalServerError)
		return
	}

	// Retorna a resposta em JSON para o cliente.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(endereco)
}

func main() {
	// Configura o endpoint /cep
	http.HandleFunc("/cep", cepHandler)

	fmt.Println("API rodando em http://localhost:8080")
	fmt.Println("Exemplo de uso: http://localhost:8080/cep?cep=01001000")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
