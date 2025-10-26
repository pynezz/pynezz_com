package main

import "time"

// InvoiceResponse represents the response for a Lightning or Monero invoice/address generation.
type InvoiceResponse struct {
	ID         string    `json:"id"`
	Method     string    `json:"method"` // "lightning" or "monero"
	Invoice    string    `json:"invoice,omitempty"` // Lightning invoice string
	Address    string    `json:"address,omitempty"` // Monero address
	PaymentID  string    `json:"payment_id,omitempty"` // Monero payment ID or subaddress
	Amount     int64     `json:"amount"` // Amount in sats or atomic units
	CreatedAt  time.Time `json:"created_at"`
	StatusURL  string    `json:"status_url"`
	QR         string    `json:"qr"` // QR code data (string, e.g. base64 or URI)
}

// StatusResponse represents the payment status for a given invoice/address.
type StatusResponse struct {
	ID        string    `json:"id"`
	Method    string    `json:"method"`
	Paid      bool      `json:"paid"`
	CheckedAt time.Time `json:"checked_at"`
}