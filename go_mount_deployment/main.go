package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// AppConfig represents the configuration for the application.
type AppConfig struct {
	ClientID      string `json:"client_id"`
	NumberOfNodes int    `json:"number_of_nodes"`
	ClientName    string `json:"client_name"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/generate-manifest", GenerateManifest).Methods("POST")

	port := "8080"
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}
}

// GenerateManifest generates a Kubernetes manifest based on the provided parameters.
func GenerateManifest(w http.ResponseWriter, r *http.Request) {
	var config AppConfig

	// Parse request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &config)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Generate manifest content
	manifestContent := fmt.Sprintf(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: %s-%s
spec:
  replicas: %d
  selector:
    matchLabels:
      app: %s-%s
  template:
    metadata:
      labels:
        app: %s-%s
    spec:
      containers:
      - name: %s-%s-container
        image: your-container-image
        ports:
        - containerPort: 80
`, config.ClientID, config.ClientName, config.NumberOfNodes, config.ClientID, config.ClientName, config.ClientID, config.ClientName, config.ClientID, config.ClientName)

	// Save manifest to a file
	manifestFileName := config.ClientID + "-" + config.ClientName + "-manifest.yaml"
	err = ioutil.WriteFile(manifestFileName, []byte(manifestContent), 0644)
	if err != nil {
		http.Error(w, "Error writing manifest to file", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Manifest file '%s' generated successfully.\n", manifestFileName)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Manifest file '%s' generated successfully.\n", manifestFileName)))
}
