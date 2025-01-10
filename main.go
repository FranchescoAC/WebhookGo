package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	Message string `json:"message"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	// Solo aceptar método POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Leer el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Parsear el JSON recibido
	var payload Payload
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Responder con un mensaje
	fmt.Printf("Received Webhook: %s\n", payload.Message)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received successfully!"))
}

func main() {
	// Configurar el endpoint del Webhook
	http.HandleFunc("/webhook", webhookHandler)

	fmt.Println("Starting server on :8080...")
	// Iniciar el servidor en el puerto 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
