package main

import (
	"fmt"
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

	for _, pre := range prefectures.Prefectures {
		fmt.Println(pre)
		db.Exec(`create table ` + pre + `(deviValue integer, schoolName text, url text)`)
	}
}
