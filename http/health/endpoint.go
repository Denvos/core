package health

import (
    "encoding/json"
    "net/http"
)

func Handler(registry *Registry, opts *Options) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        results := registry.Run(r.Context())
        status := Aggregate(results, opts)
        w.Header().Set("Content-Type", "application/json")
        if status.Status == StatusFail {
            w.WriteHeader(http.StatusServiceUnavailable)
        } else {
            w.WriteHeader(http.StatusOK)
        }
        json.NewEncoder(w).Encode(status)
    }
}

func LivenessHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{
            "status": "alive",
        })
    }
}

func ReadinessHandler(registry *Registry, opts *Options) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        results := registry.Run(r.Context())
        status := Aggregate(results, opts)
        w.Header().Set("Content-Type", "application/json")
        if status.Status == StatusFail {
            w.WriteHeader(http.StatusServiceUnavailable)
        } else {
            w.WriteHeader(http.StatusOK)
        }
        json.NewEncoder(w).Encode(status)
    }
}
