package request

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestRequest(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message":"ok"}`))
    }))
    defer server.Close()

    req := New("GET", server.URL)
    resp, err := req.Do()
    if err != nil {
        t.Fatal(err)
    }
    if !resp.IsSuccess() {
        t.Error("expected success")
    }
    var data map[string]string
    if err := resp.JSON(&data); err != nil {
        t.Fatal(err)
    }
    if data["message"] != "ok" {
        t.Errorf("expected ok, got %s", data["message"])
    }
}

func TestBuilder(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Query().Get("key") != "value" {
            t.Error("expected key=value")
        }
        w.WriteHeader(http.StatusOK)
    }))
    defer server.Close()

    builder := NewBuilder("GET", server.URL).Param("key", "value")
    req, err := builder.Build()
    if err != nil {
        t.Fatal(err)
    }
    _, err = req.Do()
    if err != nil {
        t.Fatal(err)
    }
}

func TestJSONBody(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var data map[string]string
        json.NewDecoder(r.Body).Decode(&data)
        if data["name"] != "denvos" {
            t.Errorf("expected denvos, got %s", data["name"])
        }
        w.WriteHeader(http.StatusOK)
    }))
    defer server.Close()

    req := New("POST", server.URL, WithJSONBody(map[string]string{"name": "denvos"}))
    _, err := req.Do()
    if err != nil {
        t.Fatal(err)
    }
}
