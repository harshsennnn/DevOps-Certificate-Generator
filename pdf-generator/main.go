package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"
    "github.com/jung-kurt/gofpdf"
)

// CertificateRequest is input payload for PDF generation
type CertificateRequest struct {
    Name      string `json:"name"`
    Company   string `json:"company"`
    Position  string `json:"position"`
    Duration  string `json:"duration"`
    StartDate string `json:"startdate"`
    EndDate   string `json:"enddate"`
}

// FieldMeta holds metadata about text field position and style on the template
type FieldMeta struct {
    X         float64 `json:"x"`
    Y         float64 `json:"y"`
    FontSize  int     `json:"font_size"`
    FontStyle string  `json:"font_style"` // E.g. "", "B", "I"
    Color     []int   `json:"color"`      // RGB as array of ints
}

// TemplateMeta holds all metadata for a template
type TemplateMeta struct {
    TemplateID      string               `json:"template_id"`
    BackgroundImage string               `json:"background_image"`
    Fields          map[string]FieldMeta `json:"fields"`
}

// loadTemplateMetadata loads JSON metadata for the given company/templateID
func loadTemplateMetadata(templateID string) (*TemplateMeta, error) {
    path := filepath.Join("templates", templateID+".json")
    f, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("could not open template metadata file: %w", err)
    }
    defer f.Close()

    var meta TemplateMeta
    if err := json.NewDecoder(f).Decode(&meta); err != nil {
        return nil, fmt.Errorf("failed to decode template metadata JSON: %w", err)
    }
    return &meta, nil
}

// Helper to detect image type from filename extension
func detectImageType(filename string) string {
    filename = strings.ToLower(filename)
    if strings.HasSuffix(filename, ".jpg") || strings.HasSuffix(filename, ".jpeg") {
        return "JPG"
    }
    if strings.HasSuffix(filename, ".png") {
        return "PNG"
    }
    return ""
}

func generateCertificatePDF(req CertificateRequest) ([]byte, error) {

   fullLine := fmt.Sprintf(
    "for successfully completing a %s internship from %s \n" "to %s in %s with outstanding remarks at Prodigy \n" "InfoTech.",
    req.Position,
    req.StartDate,
    req.EndDate,
    req.Duration,
)

    values := map[string]string{
        "name":          req.Name,
        "position":      req.Position,
        "duration":      req.Duration,
        "startdate":     req.StartDate,
        "enddate":       req.EndDate,
        "download_date": time.Now().Format("02/01/2006"),
        "description":   fullLine, // <-- keep this
    }

    meta, err := loadTemplateMetadata(req.Company)
    if err != nil {
        return nil, err
    }

    pdf := gofpdf.New("L", "mm", "A4", "")
    pdf.AddPage()

    pdf.ImageOptions(
        meta.BackgroundImage,
        0, 0,
        297, 210,
        false,
        gofpdf.ImageOptions{ImageType: detectImageType(meta.BackgroundImage), ReadDpi: true},
        0,
        "",
    )

    for field, pos := range meta.Fields {
        value, exists := values[field]
        if !exists || strings.TrimSpace(value) == "" {
            continue
        }

        pdf.SetFont("Arial", pos.FontStyle, float64(pos.FontSize))
        if len(pos.Color) == 3 {
            pdf.SetTextColor(pos.Color[0], pos.Color[1], pos.Color[2])
        } else {
            pdf.SetTextColor(0, 0, 0)
        }

        pdf.SetXY(pos.X, pos.Y)
        pdf.CellFormat(297, float64(pos.FontSize)+5, value, "", 0, "C", false, 0, "")
    }

    var buf bytes.Buffer
    err = pdf.Output(&buf)
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
        http.Error(w, "invalid JSON payload", http.StatusBadRequest)
        return
    }

    if req.Name == "" || req.Company == "" {
        http.Error(w, "fields 'name' and 'company' are required", http.StatusBadRequest)
        return
    }

    pdfBytes, err := generateCertificatePDF(req)
    if err != nil {
        http.Error(w, "failed to generate pdf: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", "attachment; filename=certificate.pdf")
    w.Write(pdfBytes)
}

func main() {
    http.HandleFunc("/generate", generateHandler)
    http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

    log.Println("PDF Generation Service running on :8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
