package otel

import (
	"context"
	"io"

	"go.opentelemetry.io/otel/trace"
)

type Otel struct {
	w     io.Writer
	tracer trace.Tracer
}

func New(w io.Writer, tracer trace.Tracer) *Otel {
	return &Otel{w: w, tracer: tracer}
}

func (o *Otel) Write(p []byte) (int, error) {
	ctx := context.Background()
	_, span := o.tracer.Start(ctx, "log")
	defer span.End()
	return o.w.Write(p)
}
