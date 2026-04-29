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

type EP struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Songs	   []Song `json:"songs"`
}

type Song struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Singers    []string `json:"singers"`
}

var db *sql.DB

func main() {}