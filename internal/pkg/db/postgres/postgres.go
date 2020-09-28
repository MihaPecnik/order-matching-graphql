package postgres

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func NewDatabase() {
	conn, err := gorm.Open(postgres.
		Open("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"), &gorm.Config{})

	if err != nil {
		log.Fatal("database connection error")
	}
	log.Println("successfully created database connection")

	err = conn.AutoMigrate(&Table{})
	if err != nil {
		log.Println(err.Error())
	}

	Db = conn
}

type Table struct {
	ID       int64   `json:"id"`
	UserId   int64   `json:"user_id"`
	Value    float64 `json:"value,string" sql:"type:decimal(10,2);"`
	Quantity int64   `json:"quantity"`
	Buy      bool    `json:"buy"`
	Ticker   string  `json:"ticker"`
}
