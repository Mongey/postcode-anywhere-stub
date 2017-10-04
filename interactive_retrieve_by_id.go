package main

import (
	"encoding/json"
	"net/http"
)

func InteractiveRetrieveByIDHandler(w http.ResponseWriter, r *http.Request) {
	postCode := r.URL.Query().Get("postcode")

	addr := AddressResp{
		Postcode: postCode,
	}

	resp := InteractiveResponse{
		Table: Table{
			Columns: ColumnWrapper{Column: []Column{Column{Name: "Success"}}},
			Rows: RowWrapper{
				Row: addr,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "pca-source")
	json.NewEncoder(w).Encode(resp)
}
