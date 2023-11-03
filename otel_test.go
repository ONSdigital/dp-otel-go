package dpotelgo

import "testing"
import "context"
import "time"

func TestSetup(t *testing.T){

	var cfg Config
	cfg.otel_service_name="testservice"
	cfg.otel_batch_timeout=time.Second
	cfg.otel_exporter_otlp_endpoint="localhost:4317"

	shutdown, err := SetupOTelSDK(context.Background(), cfg)

	if (shutdown==nil) {
		t.Errorf("shutdown callback is null")
	}
	if (err != nil){
		t.Errorf("err is not null %s",err)
	}

}