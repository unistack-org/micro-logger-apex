package apex

import (
	"context"
	"testing"

	log "github.com/unistack-org/micro/v3/logger"
)

func TestName(t *testing.T) {
	l2 := NewLogger(WithTextHandler())
	l2.Init()
	if l2.String() != "apex" {
		t.Errorf("name is error %s", l2.String())
	}
	t.Logf("test logger name: %s", l2.String())
}

func testLog(l log.Logger) {
	l.Infof(context.TODO(), "Test Logf with level: %s", "info")
	l.Debugf(context.TODO(), "Test Logf with level: %s", "debug")
	l.Errorf(context.TODO(), "Test Logf with level: %s", "error")
	l.Tracef(context.TODO(), "Test Logf with level: %s", "trace")
	l.Warnf(context.TODO(), "Test Logf with level: %s", "warn")
}

func TestJSON(t *testing.T) {
	l2 := NewLogger(WithJSONHandler(), WithLevel(log.TraceLevel)).Fields(map[string]interface{}{
		"Format": "JSON",
	})
	l2.Init()
	testLog(l2)
}

func TestText(t *testing.T) {
	l2 := NewLogger(WithTextHandler(), WithLevel(log.TraceLevel)).Fields(map[string]interface{}{
		"Format": "Text",
	})
	l2.Init()
	testLog(l2)
}

func TestCLI(t *testing.T) {
	l2 := NewLogger(WithCLIHandler(), WithLevel(log.TraceLevel)).Fields(map[string]interface{}{
		"Format": "CLI",
	})
	l2.Init()
	testLog(l2)
}

func TestWithLevel(t *testing.T) {
	l2 := NewLogger(WithTextHandler(), WithLevel(log.DebugLevel))
	l2.Init()
	l2.Debugf(context.TODO(), "test show debug: %s", "debug msg")

	l3 := NewLogger(WithTextHandler(), WithLevel(log.InfoLevel))
	l3.Init()
	l3.Debugf(context.TODO(), "test non-show debug: %s", "debug msg")
}

func TestWithFields(t *testing.T) {
	l2 := NewLogger(WithTextHandler()).Fields(map[string]interface{}{
		"k1": "v1",
		"k2": 123456,
	})
	l2.Init()
	l2.Info(context.TODO(), "Testing with values")
}
