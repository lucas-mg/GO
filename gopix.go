package main

import (
	"fmt"
	"github.com/bcbpix/bcbpix-go"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/pix", handlePixTransaction)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePixTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar a solicitação JSON
	var request bcbpix.PaymentRequest
	err := bcbpix.DecodeRequest(r.Body, &request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Configurar seu Certificado SSL/TLS e Chave Privada (necessário para autenticação mútua)
	certFile := "seu_certificado.crt"
	keyFile := "sua_chave_privada.key"

	client, err := bcbpix.NewClient(certFile, keyFile)
	if err != nil {
		http.Error(w, "Failed to create BCBPix client", http.StatusInternalServerError)
		return
	}

	// Enviar a transação Pix para o BCB
	response, err := client.CreatePayment(&request)
	if err != nil {
		http.Error(w, "Failed to create Pix payment", http.StatusInternalServerError)
		return
	}

	// Codificar a resposta JSON
	err = bcbpix.EncodeResponse(w, response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
