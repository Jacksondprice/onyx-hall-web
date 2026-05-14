package handlers

import (
	"encoding/json"
	"net/http"
)

type Vendor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListVendors(w http.ResponseWriter, _ *http.Request) {
	items := []Vendor{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
