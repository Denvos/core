package health

import (
    "encoding/json"
    "time"
)

type OverallStatus struct {
    Status      Status                 `json:"status"`
    Timestamp   time.Time              `json:"timestamp"`
    Details     map[string]*Result     `json:"details,omitempty"`
    Version     string                 `json:"version,omitempty"`
    Uptime      string                 `json:"uptime,omitempty"`
}

type Options struct {
    Version string
    Uptime  time.Duration
}

func Aggregate(results map[string]*Result, opts *Options) *OverallStatus {
    if opts == nil {
        opts = &Options{}
    }
    overall := &OverallStatus{
        Timestamp: time.Now(),
        Details:   results,
        Version:   opts.Version,
    }
    if opts.Uptime > 0 {
        overall.Uptime = opts.Uptime.String()
    }

    overall.Status = StatusPass
    for _, r := range results {
        if r == nil {
            continue
        }
        if r.Status == StatusFail {
            overall.Status = StatusFail
            break
        }
        if r.Status == StatusWarn && overall.Status == StatusPass {
            overall.Status = StatusWarn
        }
    }
    return overall
}

func (o *OverallStatus) MarshalJSON() ([]byte, error) {
    type Alias OverallStatus
    return json.Marshal(&struct {
        *Alias
        Details map[string]*Result `json:"details,omitempty"`
    }{
        Alias:   (*Alias)(o),
        Details: o.Details,
    })
}
