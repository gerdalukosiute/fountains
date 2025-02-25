package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gerdalukosiute/fountains/internal/database"
	"github.com/gerdalukosiute/fountains/internal/models"
	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SetupRoutes(router *mux.Router, db *sql.DB) {
	router.Use(corsMiddleware)
	router.HandleFunc("/fountains", createFountainHandler(db)).Methods("POST", "OPTIONS")
	router.HandleFunc("/fountains", listFountainsHandler(db)).Methods("GET", "OPTIONS")
	router.HandleFunc("/fountains/{id}", updateFountainHandler(db)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/fountains/{id}", deleteFountainHandler(db)).Methods("DELETE", "OPTIONS")
}

func createFountainHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fountain models.Fountain
		if err := json.NewDecoder(r.Body).Decode(&fountain); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := database.CreateFountain(db, fountain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fountain.ID = id
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(fountain)
	}
}

func listFountainsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fountains, err := database.ListFountains(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(fountains)
	}
}

func updateFountainHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid fountain ID", http.StatusBadRequest)
			return
		}

		var fountain models.Fountain
		if err := json.NewDecoder(r.Body).Decode(&fountain); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := database.UpdateFountain(db, id, fountain); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fountain.ID = int64(id)
		json.NewEncoder(w).Encode(fountain)
	}
}

func deleteFountainHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid fountain ID", http.StatusBadRequest)
			return
		}

		if err := database.DeleteFountain(db, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
