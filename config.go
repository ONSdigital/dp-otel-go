package dpotelgo

import (
	"time"
)

// Config holds the config used to initialise the OpenTelemetry Client
type Config struct {
	otelServiceName           	string
	otelExporterOtlpEndpoint 	string
	otelBatchTimeout 	 		time.Duration
}