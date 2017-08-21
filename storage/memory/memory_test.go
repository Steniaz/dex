package memory

import (
	"os"
	"testing"

	"github.com/liquidlabs-co/dex/storage"
	"github.com/liquidlabs-co/dex/storage/conformance"
	"github.com/sirupsen/logrus"
)

func TestStorage(t *testing.T) {
	logger := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.TextFormatter{DisableColors: true},
		Level:     logrus.DebugLevel,
	}

	newStorage := func() storage.Storage {
		return New(logger)
	}
	conformance.RunTests(t, newStorage)
}
