package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	_ "github.com/mattn/go-sqlite3"
)

type Saga struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Songs	   []Song `json:"songs"`
}

type Song struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Singers    	string `json:"singers"`
}

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("sqlite3", "./data/epic.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	ddl, err := os.ReadFile("./data/ddl.sql")
	if err != nil {
		log.Fatalf("Error reading DDL file: %v", err)
	}
	_, err = db.Exec(string(ddl))
	if err != nil {
		log.Fatalf("Error executing DDL: %v", err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM eps").Scan(&count)
	if err != nil {
		log.Fatalf("Error checking database: %v", err)
	}

	if count == 0 {
		log.Println("Database is empty, seeding data...")
		dml, err := os.ReadFile("./data/dml.sql")
		if err != nil {
			log.Fatalf("Error reading DML file: %v", err)
		}
		_, err = db.Exec(string(dml))
		if err != nil {
			log.Fatalf("Error executing DML: %v", err)
		}
	}
}

func main() {}


func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}