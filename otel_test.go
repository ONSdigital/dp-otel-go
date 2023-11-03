package dpotelgo

import "testing"
import "context"
import "time"

func TestSetup(t *testing.T){

	var cfg Config
	cfg.otelServiceName="testservice"
	cfg.otelBatchTimeout=time.Second
	cfg.otelExporterOtlpEndpoint="localhost:4317"

	shutdown, err := SetupOTelSDK(context.Background(), cfg)

	if (shutdown==nil) {
		t.Errorf("shutdown callback is null")
	}
	if (err != nil){
		t.Errorf("err is not null %s",err)
	}

}