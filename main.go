package main

import "github.com/jinzhu/gorm"

func main() {
	db, err := gorm.Open("postgres", "user=prefectures dbname=prefecutures sslmode=disable")
	defer db.Close()
}
