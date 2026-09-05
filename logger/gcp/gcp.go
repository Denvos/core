package gcp

import (
	"cloud.google.com/go/logging"
	"context"
)

type GCP struct {
	logger *logging.Logger
}

func New(projectID, logName string) (*GCP, error) {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	logger := client.Logger(logName)
	return &GCP{logger: logger}, nil
}

func (g *GCP) Write(p []byte) (int, error) {
	g.logger.Log(logging.Entry{Payload: string(p)})
	return len(p), nil
}

func (g *GCP) Close() error {
	return g.logger.Flush()
}
