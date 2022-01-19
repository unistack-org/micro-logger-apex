package apex

import (
	"bytes"
	"context"
	"testing"

	"go.unistack.org/micro/v3/logger"
)

func TestFields(t *testing.T) {
	ctx := context.TODO()
	buf := bytes.NewBuffer(nil)
	l := NewLogger(WithLevel(logger.TraceLevel), logger.WithOutput(buf))
	if err := l.Init(); err != nil {
		t.Fatal(err)
	}
	l.Fields("key", "val").Info(ctx, "message")
	if !bytes.Contains(buf.Bytes(), []byte(`"key":"val"`)) {
		t.Fatalf("logger fields not works, buf contains: %s", buf.Bytes())
	}
}

func TestOutput(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	l := NewLogger(logger.WithOutput(buf))
	if err := l.Init(); err != nil {
		t.Fatal(err)
	}
	l.Infof(context.TODO(), "test logger name: %s", "name")
	if !bytes.Contains(buf.Bytes(), []byte(`test logger name`)) {
		t.Fatalf("log not redirected: %s", buf.Bytes())
	}
}

func TestName(t *testing.T) {
	l2 := NewLogger(WithTextHandler())
	l2.Init()
	if l2.String() != "apex" {
		t.Errorf("name is error %s", l2.String())
	}
	t.Logf("test logger name: %s", l2.String())
}

func testLog(l logger.Logger) {
	l.Infof(context.TODO(), "Test Logf with level: %s", "info")
	l.Debugf(context.TODO(), "Test Logf with level: %s", "debug")
	l.Errorf(context.TODO(), "Test Logf with level: %s", "error")
	l.Tracef(context.TODO(), "Test Logf with level: %s", "trace")
	l.Warnf(context.TODO(), "Test Logf with level: %s", "warn")
}

func TestJSON(t *testing.T) {
	l2 := NewLogger(WithJSONHandler(), WithLevel(logger.TraceLevel)).Fields("Format", "JSON")
	l2.Init()
	testLog(l2)
}

func TestText(t *testing.T) {
	l2 := NewLogger(WithTextHandler(), WithLevel(logger.TraceLevel)).Fields("Format", "Text")
	l2.Init()
	testLog(l2)
}

func TestCLI(t *testing.T) {
	l2 := NewLogger(WithCLIHandler(), WithLevel(logger.TraceLevel)).Fields("Format", "CLI")
	l2.Init()
	testLog(l2)
}

func TestWithLevel(t *testing.T) {
	l2 := NewLogger(WithTextHandler(), WithLevel(logger.DebugLevel))
	l2.Init()
	l2.Debugf(context.TODO(), "test show debug: %s", "debug msg")

	l3 := NewLogger(WithTextHandler(), WithLevel(logger.InfoLevel))
	l3.Init()
	l3.Debugf(context.TODO(), "test non-show debug: %s", "debug msg")
}

func TestWithFields(t *testing.T) {
	l2 := NewLogger(WithTextHandler()).Fields("k1", "v1", "k2", 123456)
	l2.Init()
	l2.Info(context.TODO(), "Testing with values")
}
