package main

import (
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func main() {
	dsn := "host=localhost user=guofeiyang password= dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	type Event struct {
		DistinctId string          `gorm:"column:distinct_id" json:"distinct_id"`
		UserId     int64           `gorm:"column:user_id" json:"user_id"`
		Time       time.Time       `gorm:"column:time" json:"time"`
		Type       string          `gorm:"column:type" json:"type"`
		Event      string          `gorm:"column:event" json:"event"`
		Project    string          `gorm:"column:project" json:"project"`
		TimeFree   bool            `gorm:"column:time_free" json:"time_free"`
		Properties json.RawMessage `gorm:"column:properties" json:"properties"`
	}
	err = db.AutoMigrate(&Event{})
	if err != nil {
		panic(err)
	}
}
