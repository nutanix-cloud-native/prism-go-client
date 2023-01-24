package v3

import (
	"os"
	"testing"

	"github.com/keploy/go-sdk/kserver"
)

func TestMain(m *testing.M) {
	kserver.StartAsync()
	exitVal := m.Run()
	os.Exit(exitVal)
}
