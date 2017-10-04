package main

import (
	"encoding/json"
	"net/http"

	"github.com/icrowley/fake"
)

type BankList struct {
	Items []BankAccountResponse
}

type BankAccountResponse struct {
	IsCorrect               bool   `json:"is_correct"`
	IsDirectDebitCapable    bool   `json:"is_direct_debit_capable"`
	StatusInformation       string `json:"status_information"`
	CorrectedSortCode       string `json:"corrected_sort_code"`
	CorrectedAccountNumber  string `json:"corrected_account_number"`
	IBAN                    string `json:"iban"`
	Bank                    string `json:"bank"`
	BankBIC                 string `json:"bank_bic"`
	Branch                  string `json:"branch"`
	BranchBIC               string `json:"branch_bic"`
	ContactAddressLine1     string `json:"contact_address_line1"`
	ContactAddressLine2     string `json:"contact_address_line2"`
	ContactPostTown         string `json:"contact_post_town"`
	ContactPostcode         string `json:"contact_postcode"`
	ContactPhone            string `json:"contact_phone"`
	ContactFax              string `json:"contact_fax"`
	FasterPaymentsSupported bool   `json:"faster_payments_supported"`
	ChapsSupported          bool   `json:"chaps_supported"`
}

func BankAccountHandler(w http.ResponseWriter, r *http.Request) {
	sortCode := r.URL.Query().Get("SortCode")
	accountNumber := r.URL.Query().Get("AccountNumber")

	response := BankAccountResponse{
		IsCorrect:               true,
		IsDirectDebitCapable:    true,
		CorrectedSortCode:       sortCode,
		CorrectedAccountNumber:  accountNumber,
		ContactAddressLine1:     fake.StreetAddress(),
		ContactAddressLine2:     fake.City(),
		ContactPhone:            fake.Phone(),
		ContactFax:              fake.Phone(),
		ChapsSupported:          true,
		FasterPaymentsSupported: true,
	}
	items := BankList{
		Items: []BankAccountResponse{response},
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "pca-source")
	json.NewEncoder(w).Encode(items)
}
