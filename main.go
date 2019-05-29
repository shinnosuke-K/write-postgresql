package main

import (
	"log"
	"wite-postgresql/prefectures"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "user=prefectures password=pre dbname=prefectures sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	// create Table
	for _, pre := range prefectures.Prefectures {
		db.Exec(`create table ` + pre + `(deviValue integer, schoolName text, url text)`)
	}

	// delete Table
	//for _, pre := range prefectures.Prefectures {
	//	db.Exec(`drop table ` + pre)
	//}

}
