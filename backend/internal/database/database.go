package database

import (
	"database/sql"

	"github.com/gerdalukosiute/fountains/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS fountains (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			state TEXT CHECK(state IN ('good', 'faulty')),
			latitude REAL,
			longitude REAL
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateFountain(db *sql.DB, fountain models.Fountain) (int64, error) {
	result, err := db.Exec("INSERT INTO fountains (state, latitude, longitude) VALUES (?, ?, ?)",
		fountain.State, fountain.Latitude, fountain.Longitude)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func ListFountains(db *sql.DB) ([]models.Fountain, error) {
	rows, err := db.Query("SELECT id, state, latitude, longitude FROM fountains")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fountains []models.Fountain
	for rows.Next() {
		var f models.Fountain
		if err := rows.Scan(&f.ID, &f.State, &f.Latitude, &f.Longitude); err != nil {
			return nil, err
		}
		fountains = append(fountains, f)
	}
	return fountains, nil
}

func UpdateFountain(db *sql.DB, id int, fountain models.Fountain) error {
	_, err := db.Exec("UPDATE fountains SET state = ?, latitude = ?, longitude = ? WHERE id = ?",
		fountain.State, fountain.Latitude, fountain.Longitude, id)
	return err
}

func DeleteFountain(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM fountains WHERE id = ?", id)
	return err
}
