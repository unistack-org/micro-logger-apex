package apex

import (
	"testing"

	log "github.com/unistack-org/micro/v3/logger"
)

func TestName(t *testing.T) {
	l2 := NewLogger(WithTextHandler())
	if l2.String() != "apex" {
		t.Errorf("name is error %s", l2.String())
	}
	t.Logf("test logger name: %s", l2.String())
}

func testLog(l log.Logger) {
	l.Infof("Test Logf with level: %s", "info")
	l.Debugf("Test Logf with level: %s", "debug")
	l.Errorf("Test Logf with level: %s", "error")
	l.Tracef("Test Logf with level: %s", "trace")
	l.Warnf("Test Logf with level: %s", "warn")
}

func TestJSON(t *testing.T) {
	l2 := NewLogger(WithJSONHandler(), WithLevel(log.TraceLevel)).Fields(map[string]interface{}{
		"Format": "JSON",
	})
	testLog(l2)
}

func TestText(t *testing.T) {
	l2 := NewLogger(WithTextHandler(), WithLevel(log.TraceLevel)).Fields(map[string]interface{}{
		"Format": "Text",
	})
	testLog(l2)
}

func TestCLI(t *testing.T) {
	l2 := NewLogger(WithCLIHandler(), WithLevel(log.TraceLevel)).Fields(map[string]interface{}{
		"Format": "CLI",
	})
	testLog(l2)
}

func TestWithLevel(t *testing.T) {
	l2 := NewLogger(WithTextHandler(), WithLevel(log.DebugLevel))
	l2.Debugf("test show debug: %s", "debug msg")

	l3 := NewLogger(WithTextHandler(), WithLevel(log.InfoLevel))
	l3.Debugf("test non-show debug: %s", "debug msg")
}

func TestWithFields(t *testing.T) {
	l2 := NewLogger(WithTextHandler()).Fields(map[string]interface{}{
		"k1": "v1",
		"k2": 123456,
	})
	l2.Info("Testing with values")
}
