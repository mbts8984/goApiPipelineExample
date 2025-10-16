package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHelloHandler_DefaultName(t *testing.T) {
	// // Ensure HELLO_NAME is unset
	// os.Unsetenv("HELLO_NAME")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	helloHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	body, _ := io.ReadAll(rr.Body)
	expected := "Hello World"
	if string(body) != expected {
		t.Fatalf("expected body %q, got %q", expected, string(body))
	}
}

func TestHelloHandler_CustomName(t *testing.T) {
	os.Setenv("HELLO_NAME", "Go Developer")
	defer os.Unsetenv("HELLO_NAME")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	helloHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	body, _ := io.ReadAll(rr.Body)
	expected := "Hello Go Developer"
	if string(body) != expected {
		t.Fatalf("expected body %q, got %q", expected, string(body))
	}
}

func TestHelloHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rr := httptest.NewRecorder()

	helloHandler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405, got %d", rr.Code)
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()

	healthHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	body, _ := io.ReadAll(rr.Body)
	expected := "ok"
	if string(body) != expected {
		t.Fatalf("expected body %q, got %q", expected, string(body))
	}
}
