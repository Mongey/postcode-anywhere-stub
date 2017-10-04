package main

import (
	"flag"
	"log"
	"net/http"
)

var port, cert, key string

func init() {
	flag.StringVar(&port, "port", "9040", "Port number to bind onto")
	flag.StringVar(&cert, "cert", "server.pem", "Server certificate")
	flag.StringVar(&key, "key", "server.key", "Server key")
	flag.Parse()
}

func main() {
	http.HandleFunc("/Capture/Interactive/Find/v1.00/json3ex.ws", FindHandler)
	http.HandleFunc("/Capture/Interactive/Retrieve/v1.00/json3ex.ws", RetrieveHandler)
	http.HandleFunc("/BankAccountValidation/Interactive/Validate/v2.00/json3.ws", BankAccountHandler)
	http.HandleFunc("/PostcodeAnywhere/Interactive/Find/v1.10/xmla.ws", InteractiveFindHander)
	http.HandleFunc("/PostcodeAnywhere/Interactive/RetrieveById/v1.20/xmla.ws", InteractiveRetrieveByIDHandler)
	http.HandleFunc("/PostcodeAnywhere/Interactive/RetrieveByParts/v1.00/xmla.ws", InteractiveRetrieveByPartsHandler)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
