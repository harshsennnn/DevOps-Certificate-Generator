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
	X          float64  `json:"x"`
	Y          float64  `json:"y"`
	FontSize   int      `json:"font_size"`
	FontFamily string   `json:"font_family"`
	FontStyle  string   `json:"font_style"`
	Color      [3]int   `json:"color"`
}

type CertificateTemplate struct {
	TemplateID      string                       `json:"template_id"`
	BackgroundImage string                       `json:"background_image"`
	Fields          map[string]TextFieldPosition `json:"fields"`
}

func loadTemplate(company string) (CertificateTemplate, error) {
	var template CertificateTemplate
	path := filepath.Join("templates", company+".json")

	data, err := os.ReadFile(path)
	if err != nil {
		return template, fmt.Errorf("failed to read template JSON: %v", err)
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
	template, err := loadTemplate(req.Company)
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	// Background image
	bgPath := filepath.Join("templates", template.BackgroundImage)
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

	// Values to fill
	values := map[string]string{
		"name":          strings.ToUpper(req.Name),
		"download_date": req.DownloadDate,
		"description": fmt.Sprintf(
			"for successfully completing a %s internship from %s to %s in %s with outstanding remarks at Prodigy Infotech.",
			req.Position, req.StartDate, req.EndDate, req.Duration,
		),
	}

	// Place text
	for field, pos := range template.Fields {
		val, exists := values[field]
		if !exists || strings.TrimSpace(val) == "" {
			continue
		}

		log.Printf("Placing %q at X=%.2f, Y=%.2f", field, pos.X, pos.Y)

		if field == "description" {
			drawText(pdf, pos.X, pos.Y, pos.FontFamily, pos.FontStyle, pos.FontSize, pos.Color, val, 180)
		} else if field == "name" {
			// Center align name
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
		return nil, err
	}
	return buf.Bytes(), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var req CertificateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	pdfBytes, err := generateCertificatePDF(req)
	if err != nil {
		http.Error(w, "Failed to generate certificate: "+err.Error(), http.StatusInternalServerError)
		log.Println("PDF generation error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=certificate.pdf")
	w.Write(pdfBytes)
}

func main() {
	http.HandleFunc("/generate-certificate", handler)
	fmt.Println("Server started on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
