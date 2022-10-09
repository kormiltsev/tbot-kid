package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Dt struct {
	ChatID   string
	Perc     string
	Type     string
	Date     time.Time
	Duration time.Duration
}

const catalogFilename = "kidtracker.json"

var catalog = make([]Dt, 0)

func GetData(tipo string, date1, date2 time.Time) string {
	msg := ""
	for _, dt := range catalog {
		if date1.Before(dt.Date) && dt.Date.Before(date2) {
			if tipo == dt.Type || tipo == "none" {
				msg = msg + fmt.Sprintf("%s [%s] %s (%s)\n", dt.Date.Format("2006-01-02 15:04:05"), dt.Type, dt.Duration, dt.ChatID)
			}
		}
	}
	return msg

}

func GetFileAdres() string {
	return catalogFilename
}

func AddRow(data Dt) string {
	catalog = append(catalog, data)
	RewriteStorage()
	return "saved"
}

func GetJsonCatalog() {
	f, err := os.OpenFile(catalogFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	//считываем из файла json
	rawDataIn, err := ioutil.ReadFile(catalogFilename)
	if err != nil {
		log.Println("Cannot load catalog:", err)
	}

	err = json.Unmarshal(rawDataIn, &catalog)
	if err != nil {
		log.Println("Invalid catalogs format:", err)
	}
}

func RewriteStorage() error {
	rawDataOut, err := json.MarshalIndent(&catalog, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}

	err = ioutil.WriteFile(catalogFilename, rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated catalog file:", err)
	}
	return nil
}

func init() {
	GetJsonCatalog()
}
