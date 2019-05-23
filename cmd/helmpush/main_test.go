package main

import (
	"fmt"
	hub "github.com/Cheneytt/helm-push/pkg/nexus"
	"testing"
)

var (
	testTarballPath = "../../testdata/charts/mychart/mychart-0.1.0.tgz"
)

func TestPushCmd(t *testing.T) {

	client := hub.NewClient(
		hub.URL("http://localhost:8081/repository/helm-host/"),
		hub.Username("yourusername"),
		hub.Password("yourpassword"),
	)
	resp, err := client.UploadChartPackage(testTarballPath, false)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(resp.Status)

	err1 := handlePushResponse(resp)

	if err1 != nil {
		t.Error(err1)
	}
}
