
package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

// Structure to hold results
type DomainResult struct {
	Domain      string `json:"domain"`
	HasMX       bool   `json:"hasMX"`
	HasSPF      bool   `json:"hasSPF"`
	SPFRecord   string `json:"spfRecord"`
	HasDMARC    bool   `json:"hasDMARC"`
	DMARCRecord string `json:"dmarcRecord"`
}

// API handler
func checkDomainHandler(w http.ResponseWriter, r *http.Request) {

	//to connect to frontend and send json format
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Get domain from URL
	domain := r.URL.Query().Get("domain")
	if domain == "" {
		http.Error(w, "Domain is required", http.StatusBadRequest)
		return
	}


	result := checkDomain(domain)

	// data gets converted to json and send it to user
	json.NewEncoder(w).Encode(result)
}

// Check MX, SPF, DMARC
func checkDomain(domain string) DomainResult {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// MX checking 
	mxRecords, _ := net.LookupMX(domain)
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// SPF check
	txtRecords, _ := net.LookupTXT(domain)
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// DMARC check
	dmarcRecords, _ := net.LookupTXT("_dmarc." + domain)
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	return DomainResult{
		Domain:      domain,
		HasMX:       hasMX,
		HasSPF:      hasSPF,
		SPFRecord:   spfRecord,
		HasDMARC:    hasDMARC,
		DMARCRecord: dmarcRecord,
	}
}

func main() {
	// Serve the index.html file when visiting "/"
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// API route for domain checking
	http.HandleFunc("/check", checkDomainHandler)

	// Start server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
