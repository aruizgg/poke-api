package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aruizgg/poke-api/app.go"
)

func TestRootEndpoint(t *testing.T) {
	// Usamos la misma función setupRouter para evitar duplicación.
	router := app.SetupRouter()
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Se esperaba el status %d pero se obtuvo %d", http.StatusOK, w.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error al parsear JSON: %v", err)
	}

	expectedMessage := "Hola, mundo!"
	if response["message"] != expectedMessage {
		t.Fatalf("Se esperaba el mensaje '%s', pero se obtuvo '%s'", expectedMessage, response["message"])
	}
}
