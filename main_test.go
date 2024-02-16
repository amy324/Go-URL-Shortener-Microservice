package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gorilla/mux"
)

func TestShortenURLHandler(t *testing.T) {
    // Create a sample request body
    requestBody := `{"url": "https://example.com"}`

    // Create a request with the sample request body
    req, err := http.NewRequest("POST", "/shorten", strings.NewReader(requestBody))
    if err != nil {
        t.Fatal(err)
    }

    // Create a response recorder to record the response
    rr := httptest.NewRecorder()

    // Call the handler function directly (without running the server)
    shortenURLHandler(rr, req)

    // Check the status code of the response
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body
    expectedResponseBody := `{"short_url":"c984d06a","long_url":"https://example.com"}`
    if rr.Body.String() != expectedResponseBody {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedResponseBody)
    }
}

func TestRedirectHandler(t *testing.T) {
    // Create a new request with a GET method and a URL with a short URL parameter
    req, err := http.NewRequest("GET", "/d75277cd", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a response recorder to record the response
    rr := httptest.NewRecorder()

    // Create a mock router and register the redirectHandler
    router := mux.NewRouter()
    router.HandleFunc("/{shortURL}", func(w http.ResponseWriter, r *http.Request) {
        // Call the redirectHandler directly
        redirectHandler(w, r)
    })

    // Serve the request using the mock router
    router.ServeHTTP(rr, req)

    // Check the status code of the response
    if status := rr.Code; status != http.StatusFound {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
    }

}
