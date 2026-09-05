package check

import (
    "context"
    "database/sql"
    "net"
    "time"

    "github.com/Denvos/core/http/health"
)

type DB struct {
    db *sql.DB
    name string
}

func NewDB(db *sql.DB, name string) *DB {
    return &DB{db: db, name: name}
}

func (d *DB) Name() string {
    return d.name
}

func (d *DB) Check(ctx context.Context) (*health.Result, error) {
    if err := d.db.PingContext(ctx); err != nil {
        return &health.Result{
            Status:  health.StatusFail,
            Message: "database unreachable: " + err.Error(),
            Time:    time.Now(),
        }, nil
    }
    return &health.Result{
        Status: health.StatusPass,
        Message: "database reachable",
        Time:   time.Now(),
    }, nil
}

type HTTP struct {
    url string
    name string
}

func NewHTTP(name, url string) *HTTP {
    return &HTTP{name: name, url: url}
}

func (h *HTTP) Name() string {
    return h.name
}

func (h *HTTP) Check(ctx context.Context) (*health.Result, error) {
    client := &http.Client{Timeout: 5 * time.Second}
    req, err := http.NewRequestWithContext(ctx, "GET", h.url, nil)
    if err != nil {
        return nil, err
    }
    resp, err := client.Do(req)
    if err != nil {
        return &health.Result{
            Status: health.StatusFail,
            Message: err.Error(),
            Time:   time.Now(),
        }, nil
    }
    defer resp.Body.Close()
    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        return &health.Result{
            Status: health.StatusPass,
            Message: "HTTP check ok",
            Time:   time.Now(),
        }, nil
    }
    return &health.Result{
        Status: health.StatusFail,
        Message: "HTTP status " + resp.Status,
        Time:   time.Now(),
    }, nil
}
