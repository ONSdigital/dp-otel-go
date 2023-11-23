package dpotelgo

import (
	"time"
)

// Config holds the config used to initialise the OpenTelemetry Client
type Config struct {
	OtelServiceName          string
	OtelExporterOtlpEndpoint string
	OtelBatchTimeout         time.Duration
}
