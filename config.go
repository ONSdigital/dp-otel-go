package dpotelgo

import (
	"time"
)

// Config holds the config used to initialise the Cantabular Client
type Config struct {
	otel_service_name           string
	otel_exporter_otlp_endpoint string
	otel_batch_timeout 			time.Duration
}