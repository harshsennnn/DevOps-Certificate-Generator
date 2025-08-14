package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type CertificateRequest struct {
	Name         string `json:"name"`
	Company      string `json:"company"`
	Position     string `json:"position"`
	Duration     string `json:"duration"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	DownloadDate string `json:"download_date"`
}

func forwardToPDFGenerator(w http.ResponseWriter, r *http.Request) {
	var req CertificateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	jsonData, _ := json.Marshal(req)

	// Read target URL from environment variable, default to localhost for local dev
	pdfGeneratorURL := os.Getenv("PDF_GENERATOR_URL")
	if pdfGeneratorURL == "" {
		pdfGeneratorURL = "http://localhost:8081"
	}

	resp, err := http.Post(pdfGeneratorURL+"/generate-certificate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Error connecting to PDF generator", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Forward status code and headers
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.Header().Set("Content-Disposition", resp.Header.Get("Content-Disposition"))
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/api/certificate", forwardToPDFGenerator)
	log.Println("API Gateway running on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
