package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type info struct {
	ID         int    `json:"id"`
	Deviation  string `json:"deviation"`
	SchoolName string `json:"school_name"`
	Course     string `json:"course"`
	URL        string `json:"url"`
	Prefecture string `json:"prefecture"`
}

func (info *info) WriteCSV(db *gorm.DB) {
	count := 1
	files, err := ioutil.ReadDir("csv-name-url/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if fileName := file.Name(); fileName != ".DS_Store" {
			csvFile, err := os.Open("csv-name-url/" + fileName)
			if err != nil {
				log.Fatal(err)
			}

			reader := csv.NewReader(csvFile)

			for {
				line, err := reader.Read()
				if err != nil {
					break
				}
				info.ID = count
				count++
				info.Deviation = line[0]
				info.SchoolName = line[1]
				info.Course = line[2]
				info.URL = line[3]
				info.Prefecture = strings.Replace(fileName, ".csv", "", 1)

				db.Create(&info)
			}
		}
	}
}

func main() {
	db, err := gorm.Open("postgres", "user=prefectures password=pre dbname=prefectures sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Instantiate(Singular)
	pre := info{}

	// Instantiate(Multiple)
	//pre := []info{}

	// Create Table
	//db.CreateTable(&pre)

	// Write csv to DB
	//pre.WriteCSV(db)

	// Delete Table
	//db.DropTable(&pre)

	// Select(Singular)
	//pre.ID = 355
	//fmt.Println(pre)
	//db.First(&pre)
	//fmt.Println(pre)

	// Select(multiple)

}
