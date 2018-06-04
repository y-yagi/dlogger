package dlogger

import (
	"bytes"
	"os"
	"testing"
)

func TestDebugEnvSpecified(t *testing.T) {
	os.Setenv("DEBUG", "1")

	out := new(bytes.Buffer)
	dlogger := New(out)

	dlogger.Print("Debug message")
	expected := "[DEBUG] Debug message"
	if out.String() != expected {
		t.Errorf("Expect is %q, but %q", expected, out.String())
	}

	out.Reset()

	dlogger.Printf("Debug %v %v", "message", "with format")
	expected = "[DEBUG] Debug message with format"
	if out.String() != expected {
		t.Errorf("Expect is %q, but %q", expected, out.String())
	}
}

func TestDebugEnvNotSpecified(t *testing.T) {
	os.Setenv("DEBUG", "")

	out := new(bytes.Buffer)
	dlogger := New(out)

	dlogger.Print("Debug message")
	expected := ""
	if out.String() != expected {
		t.Errorf("Expect is %q, but %q", expected, out.String())
	}

	out.Reset()

	dlogger.Printf("Debug %v %v", "message", "with format")
	expected = ""
	if out.String() != expected {
		t.Errorf("Expect is %q, but %q", expected, out.String())
	}
}

func TestMain(m *testing.M) {
	beforeEnv := os.Getenv("DEBUG")
	result := m.Run()
	os.Setenv("DEBUG", beforeEnv)
	os.Exit(result)
}
