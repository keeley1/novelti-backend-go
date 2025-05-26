package utils

import (
	"fmt"
	"testing"
)

func TestCloseBody_Success(t *testing.T) {
	m := &mockCloser{shouldFail: false}

	err := CloseBody(m)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !m.closed {
		t.Error("expected body to be closed")
	}
}

// Test that failures are handled
func TestCloseBody_Failure(t *testing.T) {
	m := &mockCloser{shouldFail: true}

	err := CloseBody(m)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if !m.closed {
		t.Error("expected body to be closed")
	}
}

// Mock closer for testing
type mockCloser struct {
	shouldFail bool
	closed     bool
}

func (m *mockCloser) Close() error {
	m.closed = true
	if m.shouldFail {
		return fmt.Errorf("close error")
	}
	return nil
}
