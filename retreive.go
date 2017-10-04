package main

import (
	"encoding/json"
	"net/http"
)

func RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	items := AddressList{
		Items: make([]Address, 0),
	}
	id := r.URL.Query().Get("Id")
	if id == "GB|RM|A|2" {
		address := Address{
			Id:                   "GB|RM|A|2",
			DomesticId:           "2",
			Language:             "ENG",
			LanguageAlternatives: "ENG",
			Company:              "Super Long Widget Company",
			SubBuilding:          "Widget Research Department",
			BuildingName:         "The Big Blue Building",
			Street:               "The Longest Street In The World",
			Line1:                "This is a line that should exceed the limit of fifty characters",
			Line2:                "Line 2",
			District:             "Central District",
			City:                 "London",
			PostalCode:           "SW1 1AA",
			CountryName:          "United Kingdom",
			Label:                "This is a really long address",
			Type:                 "Commercial",
			DataLevel:            "Premise",
		}
		items.Items = append(items.Items, address)
	} else {
		address := Address{
			Id:                   "GB|RM|A|1000676865",
			DomesticId:           "1000676865",
			Language:             "ENG",
			LanguageAlternatives: "ENG",
			Company:              "Ritz Hotel",
			Street:               "Arlington Street",
			Line1:                "Arlington Street",
			City:                 "London",
			PostalCode:           "SW1A 1RB",
			CountryName:          "United Kingdom",
			Label:                "Ritz Hotel\nArlington Street\nLONDON\nSW1A 1RB\nUNITED KINGDOM",
			Type:                 "Commercial",
			DataLevel:            "Premise",
		}
		items.Items = append(items.Items, address)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "pca-source")
	json.NewEncoder(w).Encode(items)
}
