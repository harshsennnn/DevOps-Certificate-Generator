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

