package main

import (
	"encoding/json"
	"net/http"
)

type ItemList struct {
	Items []Item
}

func FindHandler(w http.ResponseWriter, r *http.Request) {
	items := ItemList{
		Items: make([]Item, 0),
	}

	container := r.URL.Query().Get("Container")
	if container == "" {
		item := Item{
			Id:          "GB|RM|A|P-SW1A-1RB",
			Type:        "Postcode",
			Text:        "SW1A 1RB",
			Highlight:   "0-1,1-2,2-3,3-4,5-6,6-7,7-8",
			Description: "Arlington Street, London",
		}
		items.Items = append(items.Items, item)
	} else {
		item := Item{
			Id:          "GB|RM|A|1000676865",
			Type:        "Address",
			Text:        "Ritz Hotel, Arlington Street",
			Highlight:   "",
			Description: "London, SW1A 1RB",
		}
		items.Items = append(items.Items, item)

		item = Item{
			Id:          "GB|RM|A|2",
			Type:        "Address",
			Text:        "The long test",
			Highlight:   "",
			Description: "A long address",
		}
		items.Items = append(items.Items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "pca-source")
	json.NewEncoder(w).Encode(items)
}
