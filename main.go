package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Saga struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Songs       []Song `json:"songs"`
}

type Song struct {
	ID          int    `json:"id"`
	SagaID      int    `json:"saga_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Singers     string `json:"singers"`
	ImageURL    string `json:"image_url"`
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
	err = db.QueryRow("SELECT COUNT(*) FROM sagas").Scan(&count)
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

func extractID(r *http.Request, prefix string) (int, bool) {
	path := strings.TrimPrefix(r.URL.Path, prefix)
	if path == "" {
		return 0, false
	}
	id, err := strconv.Atoi(path)
	if err != nil {
		return 0, false
	}
	return id, true
}

func sagasHandler(w http.ResponseWriter, r *http.Request) {
	id, hasID := extractID(r, "/sagas/")
	if hasID {
		switch r.Method {
		case http.MethodGet:
			handleGetSagaByID(w, id)
		case http.MethodPut:
			handleUpdateSaga(w, r, id)
		case http.MethodDelete:
			handleDeleteSaga(w, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}
	switch r.Method {
	case http.MethodGet:
		handleGetSagas(w, r)
	case http.MethodPost:
		handleCreateSaga(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func songsHandler(w http.ResponseWriter, r *http.Request) {
	id, hasID := extractID(r, "/songs/")
	if hasID {
		switch r.Method {
		case http.MethodGet:
			handleGetSongByID(w, id)
		case http.MethodPut:
			handleUpdateSong(w, r, id)
		case http.MethodDelete:
			handleDeleteSong(w, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}
	switch r.Method {
	case http.MethodPost:
		handleCreateSong(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	defer db.Close()

	http.HandleFunc("/sagas", corsMiddleware(sagasHandler))
	http.HandleFunc("/sagas/", corsMiddleware(sagasHandler))
	http.HandleFunc("/songs", corsMiddleware(songsHandler))
	http.HandleFunc("/songs/", corsMiddleware(songsHandler))

	log.Println("Server is running on http://localhost:8088")
	log.Fatal(http.ListenAndServe(":8088", nil))
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func handleCreateSaga(w http.ResponseWriter, r *http.Request) {
	var s Saga
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(s.Title) == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(
		"INSERT INTO sagas (title, description, image_url) VALUES (?, ?, ?)",
		s.Title, s.Description, s.ImageURL,
	)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	s.ID = int(id)
	writeJSON(w, http.StatusCreated, s)
}

func handleGetSagas(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, description, image_url FROM sagas")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	sagas := []Saga{}
	for rows.Next() {
		var s Saga
		rows.Scan(&s.ID, &s.Title, &s.Description, &s.ImageURL)
		sagas = append(sagas, s)
	}

	writeJSON(w, http.StatusOK, sagas)
}

func handleGetSagaByID(w http.ResponseWriter, id int) {
	var s Saga
	err := db.QueryRow(
		"SELECT id, title, description, image_url FROM sagas WHERE id = ?", id,
	).Scan(&s.ID, &s.Title, &s.Description, &s.ImageURL)
	if err == sql.ErrNoRows {
		http.Error(w, "Saga not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query(
		"SELECT id, saga_id, title, description, singers, image_url FROM songs WHERE saga_id = ?", id,
	)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	s.Songs = []Song{}
	for rows.Next() {
		var song Song
		rows.Scan(&song.ID, &song.SagaID, &song.Title, &song.Description, &song.Singers, &song.ImageURL)
		s.Songs = append(s.Songs, song)
	}

	writeJSON(w, http.StatusOK, s)
}

func handleUpdateSaga(w http.ResponseWriter, r *http.Request, id int) {
	var s Saga
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(
		"UPDATE sagas SET title = ?, description = ?, image_url = ? WHERE id = ?",
		s.Title, s.Description, s.ImageURL, id,
	)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		http.Error(w, "Saga not found", http.StatusNotFound)
		return
	}

	s.ID = id
	writeJSON(w, http.StatusOK, s)
}

func handleDeleteSaga(w http.ResponseWriter, id int) {
	res, err := db.Exec("DELETE FROM sagas WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		http.Error(w, "Saga not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Saga deleted"})
}

func handleCreateSong(w http.ResponseWriter, r *http.Request) {
	var s Song
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(s.Title) == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	if s.SagaID == 0 {
		http.Error(w, "saga_id is required", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(
		"INSERT INTO songs (saga_id, title, description, singers, image_url) VALUES (?, ?, ?, ?, ?)",
		s.SagaID, s.Title, s.Description, s.Singers, s.ImageURL,
	)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	s.ID = int(id)
	writeJSON(w, http.StatusCreated, s)
}

func handleGetSongByID(w http.ResponseWriter, id int) {
	var s Song
	err := db.QueryRow(
		"SELECT id, saga_id, title, description, singers, image_url FROM songs WHERE id = ?", id,
	).Scan(&s.ID, &s.SagaID, &s.Title, &s.Description, &s.Singers, &s.ImageURL)
	if err == sql.ErrNoRows {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, s)
}

func handleUpdateSong(w http.ResponseWriter, r *http.Request, id int) {
	var s Song
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	res, err := db.Exec(
		"UPDATE songs SET title = ?, description = ?, singers = ?, image_url = ? WHERE id = ?",
		s.Title, s.Description, s.Singers, s.ImageURL, id,
	)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}

	s.ID = id
	writeJSON(w, http.StatusOK, s)
}

func handleDeleteSong(w http.ResponseWriter, id int) {
	res, err := db.Exec("DELETE FROM songs WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Song deleted"})
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }
        next(w, r)
    }
}