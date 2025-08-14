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

	"github.com/jung-kurt/gofpdf"
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

type TextFieldPosition struct {
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	FontSize   int     `json:"font_size"`
	FontFamily string  `json:"font_family"`
	FontStyle  string  `json:"font_style"`
	Color      [3]int  `json:"color"`
}

type CertificateTemplate struct {
	TemplateID      string                       `json:"template_id"`
	BackgroundImage string                       `json:"background_image"`
	Fields          map[string]TextFieldPosition `json:"fields"`
}

// CORS middleware
func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		allowedOrigins := []string{
			"http://localhost:5173",
			"http://localhost",
			"http://127.0.0.1",
		}

		for _, o := range allowedOrigins {
			if origin == o {
				w.Header().Set("Access-Control-Allow-Origin", o)
				break
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		h(w, r)
	}
}

func loadTemplate(company string) (CertificateTemplate, error) {
	var template CertificateTemplate

	// Absolute path inside container
	path := filepath.Join("./templates", company+".json")
	log.Println("Loading template from:", path)

	data, err := os.ReadFile(path)
	if err != nil {
		return template, fmt.Errorf("failed to read template JSON at %s: %v", path, err)
	}

	if err := json.Unmarshal(data, &template); err != nil {
		return template, fmt.Errorf("failed to parse template JSON: %v", err)
	}

	return template, nil
}

func drawText(pdf *gofpdf.Fpdf, x, y float64, fontFamily, fontStyle string, fontSize int, color [3]int, text string, maxWidth float64) {
	pdf.SetFont(fontFamily, fontStyle, float64(fontSize))
	pdf.SetTextColor(color[0], color[1], color[2])

	if maxWidth <= 0 {
		pdf.Text(x, y, text)
		return
	}

	lines := pdf.SplitLines([]byte(text), maxWidth)
	lineHeight := float64(fontSize) * 0.5
	for i, line := range lines {
		pdf.Text(x, y+(float64(i)*lineHeight), string(line))
	}
}

func generateCertificatePDF(req CertificateRequest) ([]byte, error) {
	log.Printf("Generating PDF for: %+v\n", req)

	template, err := loadTemplate(req.Company)
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	bgPath := filepath.Join("./templates", template.BackgroundImage)
	log.Println("Loading background image from:", bgPath)

	ext := strings.ToLower(filepath.Ext(bgPath))
	imgType := map[string]string{
		".png":  "PNG",
		".jpg":  "JPG",
		".jpeg": "JPG",
	}[ext]
	if imgType == "" {
		imgType = "PNG"
	}
	pdf.ImageOptions(bgPath, 0, 0, 297, 210, false, gofpdf.ImageOptions{ImageType: imgType, ReadDpi: true}, 0, "")

	values := map[string]string{
		"name":          strings.ToUpper(req.Name),
		"download_date": req.DownloadDate,
		"description": fmt.Sprintf(
			"for successfully completing a %s internship from %s to %s in %s with outstanding remarks at Prodigy Infotech.",
			req.Position, req.StartDate, req.EndDate, req.Duration,
		),
	}

	for field, pos := range template.Fields {
		val, exists := values[field]
		if !exists || strings.TrimSpace(val) == "" {
			continue
		}

		if field == "description" {
			drawText(pdf, pos.X, pos.Y, pos.FontFamily, pos.FontStyle, pos.FontSize, pos.Color, val, 180)
		} else if field == "name" {
			pdf.SetFont(pos.FontFamily, pos.FontStyle, float64(pos.FontSize))
			width := pdf.GetStringWidth(val)
			centerX := (310 - width) / 2
			drawText(pdf, centerX, pos.Y, pos.FontFamily, pos.FontStyle, pos.FontSize, pos.Color, val, 0)
		} else {
			drawText(pdf, pos.X, pos.Y, pos.FontFamily, pos.FontStyle, pos.FontSize, pos.Color, val, 0)
		}
	}

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, fmt.Errorf("failed to output PDF: %v", err)
	}
	return buf.Bytes(), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var req CertificateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Invalid JSON:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	pdfBytes, err := generateCertificatePDF(req)
	if err != nil {
		log.Println("Error generating PDF:", err)
		http.Error(w, "Failed to generate certificate: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=certificate.pdf")
	w.Write(pdfBytes)
}

func getTemplates(w http.ResponseWriter, r *http.Request) {
	tmplPath := "./templates/companyA.json"
	data, err := os.ReadFile(tmplPath)
	if err != nil {
		log.Println("Error reading template list:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func init() {
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)
}

func main() {
	http.HandleFunc("/generate-certificate", withCORS(handler))
	http.HandleFunc("/templates", withCORS(getTemplates))

	http.Handle("/templates/", withCORS(func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))).ServeHTTP(w, r)
	}))

	log.Println("PDF Generator running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
