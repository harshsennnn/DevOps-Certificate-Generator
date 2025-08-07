package main

import (
    "bytes"
    "encoding/json"
    "io"
    "log"
    "net/http"
)

type CertificateRequest struct {
    Name      string `json:"name"`
    Company   string `json:"company"` // e.g., template ID
    Position  string `json:"position"`
    Duration  string `json:"duration"`
    StartDate string `json:"startdate"`
    EndDate   string `json:"enddate"`
}


func generateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "only POST method allowed", http.StatusMethodNotAllowed)
        return
    }

    var req CertificateRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid JSON payload", http.StatusBadRequest)
        return
    }

    // Basic validation
    if req.Name == "" || req.Company == "" || req.Position == "" || req.Duration == "" {
        http.Error(w, "all fields are required", http.StatusBadRequest)
        return
    }

    // Marshal request for forwarding
    reqBody, err := json.Marshal(req)
    if err != nil {
        http.Error(w, "failed to marshal request", http.StatusInternalServerError)
        return
    }

    // Forward to PDF generation service
    resp, err := http.Post("http://localhost:8081/generate", "application/json", bytes.NewBuffer(reqBody))
    if err != nil {
        http.Error(w, "failed to call pdf generation service: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        http.Error(w, "pdf generation error: "+string(body), resp.StatusCode)
        return
    }

    // Copy PDF bytes to response
    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", "attachment; filename=certificate.pdf")
    io.Copy(w, resp.Body)
}

func main() {
    http.HandleFunc("/generate", generateHandler)
    log.Println("API Gateway running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
