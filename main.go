package main

import (
	"log"
	"write-postgresql/prefectures"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createTable(db *gorm.DB) {
	// create Table
	for _, pre := range prefectures.Prefectures {
		db.Exec(`create table ` + pre + `(deviValue integer, schoolName text, course text, url text)`)
	}
}

func deleteDB(db *gorm.DB) {
	//delete Table
	for _, pre := range prefectures.Prefectures {
		//db.Exec(`drop table ` + pre)
		db.DropTable(pre)
	}
}

func main() {
	db, err := gorm.Open("postgres", "user=prefectures password=pre dbname=prefectures sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	pre := info{}

	//db.AutoMigrate(&pre)

	//pre.SchoolName = "sdfafsad"
	//db.Create(&pre)

	db.First(&pre)
	db.Delete(&pre)

	//createTable(db)
	//deleteDB(db)

}
