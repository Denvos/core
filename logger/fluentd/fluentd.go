package fluentd

import (
	"encoding/json"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
)

type Fluentd struct {
	logger *fluent.Fluent
	tag    string
}

func New(host, port string, tag string) (*Fluentd, error) {
	logger, err := fluent.New(fluent.Config{
		FluentHost: host,
		FluentPort: port,
		Timeout:    5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &Fluentd{
		logger: logger,
		tag:    tag,
	}, nil
}

func (f *Fluentd) Write(p []byte) (int, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(p, &data); err != nil {
		data = map[string]interface{}{
			"message": string(p),
		}
	}
	err := f.logger.Post(f.tag, data)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (f *Fluentd) Close() error {
	return f.logger.Close()
}
