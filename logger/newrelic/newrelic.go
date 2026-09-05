package newrelic

import (
	"io"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type NewRelic struct {
	app    *newrelic.Application
	writer io.Writer
}

func New(license string, appName string) (*NewRelic, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(license),
	)
	if err != nil {
		return nil, err
	}
	return &NewRelic{app: app}, nil
}

func (n *NewRelic) Write(p []byte) (int, error) {
	txn := n.app.StartTransaction("log", nil, nil)
	defer txn.End()
	txn.NoticeError(newrelic.Error{
		Message: string(p),
		Class:   "log",
	})
	return len(p), nil
}
