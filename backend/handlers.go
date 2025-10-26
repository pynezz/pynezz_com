package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// InvoiceResponse represents the response for invoice/address generation
type InvoiceResponse struct {
	ID        string `json:"id"`
	Invoice   string `json:"invoice,omitempty"`
	Address   string `json:"address,omitempty"`
	PaymentID string `json:"payment_id,omitempty"`
	StatusURL string `json:"status_url"`
	QRString  string `json:"qr_string"`
	Method    string `json:"method"`
}

// StatusResponse represents the response for payment status
type StatusResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"` // "pending" or "paid"
	Method string `json:"method"`
}

// handleLightningInvoice handles POST /api/invoice/lightning
func handleLightningInvoice(w http.ResponseWriter, r *http.Request) {
	id := randomID()
	invoice := "lnbc1" + randomString(50)
	resp := InvoiceResponse{
		ID:        id,
		Invoice:   invoice,
		StatusURL: "/api/status/" + id,
		QRString:  invoice,
		Method:    "lightning",
	}
	saveMockStatus(id, "pending", "lightning")
	writeJSON(w, resp)
}

// handleMoneroInvoice handles POST /api/invoice/monero
func handleMoneroInvoice(w http.ResponseWriter, r *http.Request) {
	id := randomID()
	address := "4" + randomString(94)
	paymentID := randomString(16)
	resp := InvoiceResponse{
		ID:        id,
		Address:   address,
		PaymentID: paymentID,
		StatusURL: "/api/status/" + id,
		QRString:  address,
		Method:    "monero",
	}
	saveMockStatus(id, "pending", "monero")
	writeJSON(w, resp)
}

// handleStatus handles GET /api/status/{id}
func handleStatus(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/status/")
	status, method := getMockStatus(id)
	resp := StatusResponse{
		ID:     id,
		Status: status,
		Method: method,
	}
	writeJSON(w, resp)
}

// --- Mock payment status store ---

var mockStatusStore = make(map[string]struct {
	Status string
	Method string
	Time   time.Time
})

func saveMockStatus(id, status, method string) {
	mockStatusStore[id] = struct {
		Status string
		Method string
		Time   time.Time
	}{Status: status, Method: method, Time: time.Now()}
}

func getMockStatus(id string) (string, string) {
	entry, ok := mockStatusStore[id]
	if !ok {
		return "not_found", ""
	}
	// Simulate payment after 10 seconds
	if entry.Status == "pending" && time.Since(entry.Time) > 10*time.Second {
		entry.Status = "paid"
		mockStatusStore[id] = entry
	}
	return entry.Status, entry.Method
}

// --- Helpers ---

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func randomID() string {
	return randomString(12)
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
