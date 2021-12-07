package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Response struct {
	Message string `json:"message"`
}

func TestHomeHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	HomeHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	t.Log(w.Body.String())
	got := Response{}
	err := json.NewDecoder(w.Body).Decode(&got)
	if err != nil {
		t.Errorf("Error decoding response body %s", err)
	}
	want := Response{Message: "Welcome to the Q4 GO Bootcamp API"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expeted %v but got %v", want, got)
	}
}

func TestGetCharacters(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/characters", nil)
	GetCharacters(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}

func TestSaveCharacter(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/save_characters", nil)
	SaveCharacters(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}
