package datadog

import (
	"context"
	"io"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

type Datadog struct {
	client   *datadogV2.LogsApi
	ctx      context.Context
	service  string
	source   string
	hostname string
	tags     []string
}

type Option func(*Datadog)

func WithService(service string) Option {
	return func(d *Datadog) {
		d.service = service
	}
}

func WithSource(source string) Option {
	return func(d *Datadog) {
		d.source = source
	}
}

func WithHostname(hostname string) Option {
	return func(d *Datadog) {
		d.hostname = hostname
	}
}

func WithTags(tags []string) Option {
	return func(d *Datadog) {
		d.tags = tags
	}
}

func New(opts ...Option) (*Datadog, error) {
	ctx := context.Background()
	cfg := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(cfg)
	api := datadogV2.NewLogsApi(apiClient)

	d := &Datadog{
		client: api,
		ctx:    ctx,
	}
	for _, opt := range opts {
		opt(d)
	}
	return d, nil
}

func (d *Datadog) Write(p []byte) (int, error) {
	log := datadogV2.NewHTTPLogItem()
	log.SetMessage(string(p))

	if d.service != "" {
		log.SetService(d.service)
	}
	if d.source != "" {
		log.SetSource(d.source)
	}
	if d.hostname != "" {
		log.SetHostname(d.hostname)
	}
	if len(d.tags) > 0 {
		log.SetTags(d.tags)
	}

	_, _, err := d.client.SubmitLog(d.ctx, []datadogV2.HTTPLogItem{*log})
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
