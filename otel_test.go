package dpotelgo

import "testing"
import "context"
import "os"

func TestSetup(t *testing.T){

	os.Setenv("OTEL_SERVICE_NAME","testservice")
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT","testendpoint")
	shutdown, err := SetupOTelSDK(context.Background())

	if (shutdown==nil) {
		t.Errorf("shutdown callback is null")
	}
	if (err != nil){
		t.Errorf("err is not null %s",err)
	}

}