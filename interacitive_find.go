package main

import (
	"encoding/json"
	"net/http"

	"github.com/icrowley/fake"
)

type AddressListItem struct {
	ID            string `json:"Id"`
	StreetAddress string `json:"StreetAddress"`
	Place         string `json:"Place"`
}

func InteractiveFindHander(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("SearchTerm")

	resp := InteractiveResponse{
		Table: Table{
			Columns: ColumnWrapper{Column: []Column{Column{Name: "Success"}}},
			Rows: RowWrapper{
				Row: AddressListItem{
					ID:            searchTerm,
					StreetAddress: fake.StreetAddress(),
					Place:         fake.City(),
				},
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "pca-source")
	json.NewEncoder(w).Encode(resp)
	//data.parsed_response["Table"]["Rows"]["Row"].each do |item|
}
