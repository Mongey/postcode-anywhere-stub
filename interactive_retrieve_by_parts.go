package main

import (
	"encoding/json"
	"net/http"
)

type AddressResp struct {
	Mailsort          string
	Barcode           string
	Type              string
	Udprn             string
	Company           string
	Department        string
	Postcode          string
	Line1             string
	Line2             string
	Line3             string
	Line4             string
	Line5             string
	BuildingName      string
	BuildingNumber    string
	PrimaryStreet     string
	SubBuilding       string
	DependentLocality string
	PostTown          string
	County            string
}

func InteractiveRetrieveByPartsHandler(w http.ResponseWriter, r *http.Request) {
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
