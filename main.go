package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
	"powerlifting/internal/calculator"
	"powerlifting/internal/config"
)

func main() {
	defaultPath := "./internal/config/config.json"
	configPath := flag.String("c", defaultPath, "config file path")
	flag.Parse()

	configFile, err := os.Open(*configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	var cfg *config.Config

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&cfg); err != nil {
		log.Fatal(err)
		return
	}

	c := calculator.NewCalculator(cfg)

	cc, err := c.Calculate()
	if err != nil {
		log.Fatal(err)
		return
	}

	result := make([]string, 0, len(cc)+1)
	result = append(result, cfg.Maximums.String())
	result = append(result, cc...)

	err = saveAsPDF(cfg.Name, result...)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func saveAsPDF(name string, txt ...string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	pdf := gofpdf.New("P", "mm", "A4", pwd+"/font/")
	pdf.AddFont("Helvetica", "", "helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 8)

	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")

	for _, c := range txt {
		pdf.MultiCell(200, 4, tr(c+"\n"), "", "", false)

	}

	err = pdf.OutputFileAndClose(fmt.Sprintf("%v.pdf", name))
	if err != nil {
		return err
	}

	return nil
}
