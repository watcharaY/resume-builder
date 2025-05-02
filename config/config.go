package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig() (htmlPath, pdfPath string, err error) {
	err = godotenv.Load()
	if err != nil {
		return "", "", err
	}

	htmlPath = os.Getenv("HTML_PATH")
	pdfPath = os.Getenv("PDF_PATH")

	if htmlPath == "" {
		return "", "", errors.New("html path is empty")
	}
	if pdfPath == "" {
		return "", "", errors.New("pdf path is empty")
	}

	return htmlPath, pdfPath, nil
}
