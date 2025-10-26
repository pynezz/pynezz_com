package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type PaymentStatus string

const (
	StatusPending PaymentStatus = "pending"
	StatusPaid    PaymentStatus = "paid"
)

type MockInvoice struct {
	ID        string
	Method    string // "lightning" or "monero"
	Invoice   string // Lightning invoice string or Monero address
	Extra     string // Payment hash for Lightning, Payment ID for Monero
	CreatedAt time.Time
	Status    PaymentStatus
}

var (
	mockPayments   = make(map[string]*MockInvoice)
	mockPaymentsMu sync.RWMutex
)

func init() {
	go simulatePayments()
}

// Generate a random string for IDs, hashes, etc.
func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Create a mock Lightning invoice
func createMockLightningInvoice() *MockInvoice {
	id := randomString(12)
	invoice := fmt.Sprintf("lnbc1%s", randomString(40))
	paymentHash := randomString(64)
	m := &MockInvoice{
		ID:        id,
		Method:    "lightning",
		Invoice:   invoice,
		Extra:     paymentHash,
		CreatedAt: time.Now(),
		Status:    StatusPending,
	}
	mockPaymentsMu.Lock()
	mockPayments[id] = m
	mockPaymentsMu.Unlock()
	return m
}

// Create a mock Monero address + payment ID
func createMockMoneroInvoice() *MockInvoice {
	id := randomString(12)
	address := fmt.Sprintf("4%s", randomString(93))
	paymentID := randomString(16)
	m := &MockInvoice{
		ID:        id,
		Method:    "monero",
		Invoice:   address,
		Extra:     paymentID,
		CreatedAt: time.Now(),
		Status:    StatusPending,
	}
	mockPaymentsMu.Lock()
	mockPayments[id] = m
	mockPaymentsMu.Unlock()
	return m
}

// Get payment status by ID
func getMockPaymentStatus(id string) (PaymentStatus, bool) {
	mockPaymentsMu.RLock()
	defer mockPaymentsMu.RUnlock()
	m, ok := mockPayments[id]
	if !ok {
		return "", false
	}
	return m.Status, true
}

// Simulate payments being paid after a random delay
func simulatePayments() {
	for {
		time.Sleep(2 * time.Second)
		now := time.Now()
		mockPaymentsMu.Lock()
		for _, m := range mockPayments {
			if m.Status == StatusPending && now.Sub(m.CreatedAt) > time.Duration(3+rand.Intn(5))*time.Second {
				m.Status = StatusPaid
			}
		}
		mockPaymentsMu.Unlock()
	}
}
