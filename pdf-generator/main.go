package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "path/filepath"

    "github.com/jung-kurt/gofpdf"
)

// CertificateRequest represents input payload for PDF generation
type CertificateRequest struct {
    Name     string `json:"name"`
    Company  string `json:"company"`
    Position string `json:"position"`
    Duration string `json:"duration"`
}

// Predefined map from company keys to template images stored locally in ./templates/
var companyTemplates = map[string]string{
    "companyA": "templates/companyA_template.jpg",
    "companyB": "templates/companyB_template.jpg",
    "companyC": "templates/companyC_template.jpg",
    "companyD": "templates/companyD_template.jpg",
    "companyE": "templates/companyE_template.jpg",
}

func generateCertificatePDF(req CertificateRequest) ([]byte, error) {
    templatePath, ok := companyTemplates[req.Company]
    if !ok {
        return nil, fmt.Errorf("unknown company template: %s", req.Company)
    }

    pdf := gofpdf.New("L", "mm", "A4", "") 
    pdf.AddPage()

    // Add background image covering full A4 page (297x210)
    pdf.ImageOptions(
        templatePath,
        0, 0,
        297, 210,
        false,
        gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
        0,
        "",
    )

    // Overlay text fields
    pdf.SetFont("Arial", "B", 24)
    pdf.SetTextColor(0, 0, 128)
    pdf.SetXY(50, 90)
    pdf.CellFormat(200, 15, req.Name, "", 0, "C", false, 0, "")

    pdf.SetFont("Arial", "", 18)
    pdf.SetXY(50, 110)
    pdf.CellFormat(200, 12, fmt.Sprintf("Position: %s", req.Position), "", 0, "C", false, 0, "")

    pdf.SetXY(50, 125)
    pdf.CellFormat(200, 12, fmt.Sprintf("Duration: %s", req.Duration), "", 0, "C", false, 0, "")

    var buf bytes.Buffer
    err := pdf.Output(&buf)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "only POST method allowed", http.StatusMethodNotAllowed)
        return
    }

    var req CertificateRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid json payload", http.StatusBadRequest)
        return
    }

    pdfBytes, err := generateCertificatePDF(req)
    if err != nil {
        http.Error(w, "failed to generate pdf: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/pdf")
    w.Write(pdfBytes)
}

func main() {
    http.HandleFunc("/generate", generateHandler)

    // Serve the templates directory for manual testing
    http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir(filepath.Join(".", "templates")))))

    log.Println("PDF Generation Service running on :8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
