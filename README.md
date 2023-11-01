# dp-otel-go
OpenTelemetry startup code for Go services

# How to use
You need the following environment variables set:
```
    OTEL_SERVICE_NAME
	OTEL_EXPORTER_OTLP_ENDPOINT
```
NB the endpoint should be of the form: "hostname:<portnum>" - no protocol identifier

in your service initialisation code:
```
    shutdown, err := SetupOTelSDK(context.Background())
```


# TO DO
