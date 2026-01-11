package main

import (
	"encoding/json"
	"net/http"

	"minidb/internal/engine"
)

func main() {
	eng, _ := engine.NewEngine("data.wal")

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		rows, _ := eng.Execute(r.Context(), "SELECT * FROM users")
		json.NewEncoder(w).Encode(rows)
	})

	http.ListenAndServe(":8080", nil)
}
