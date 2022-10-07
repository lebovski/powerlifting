package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"calculate/internal/calculator"
	"calculate/internal/config"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	defaultPath := "./internal/config/config.json"
	configPath := flag.String("c", defaultPath, "config file path")

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

	max, err := readFromConsole()
	if err != nil {
		log.Fatal(err)
		return
	}

	c := calculator.NewCalculator(max, cfg)

	cc, err := c.Calculate()
	if err != nil {
		log.Fatal(err)
		return
	}

	result := make([]string, 0, len(cc)+1)
	result = append(result, max.String())
	result = append(result, cc...)

	err = saveAsPDF(result...)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func saveAsPDF(txt ...string) error {
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

	err = pdf.OutputFileAndClose("train.pdf")
	if err != nil {
		return err
	}

	return nil
}

func readFromConsole() (*calculator.MaxExercise, error) {
	reader := bufio.NewReader(os.Stdin)

	benchPress, err := readFromStdIn(reader, "Max Жим лёжа: ")
	if err != nil {
		return nil, err
	}

	deadLift, err := readFromStdIn(reader, "Max Становая тяга: ")
	if err != nil {
		return nil, err
	}

	squat, err := readFromStdIn(reader, "Max Присед: ")
	if err != nil {
		return nil, err
	}

	fmt.Print("\n")

	return &calculator.MaxExercise{
		BenchPress: benchPress,
		DeadLift:   deadLift,
		Squat:      squat,
	}, nil
}

func readFromStdIn(reader *bufio.Reader, msg string) (float64, error) {
	fmt.Print(msg)
	val, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	val = strings.Trim(val, " \n")

	res, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, err
	}

	return res, nil
}
