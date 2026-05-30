package main

import (
	"os"
	"testing"
)

func TestRun_Error(t *testing.T) {
	os.Setenv("APP_PORT", "invalid")
	defer os.Unsetenv("APP_PORT")

	err := Run()
	if err == nil {
		t.Error("Expected error when running with invalid port, got nil")
	}
}

func TestMain_Error(t *testing.T) {
	os.Setenv("APP_PORT", "invalid")
	defer os.Unsetenv("APP_PORT")

	// Mock logFatal
	origLogFatal := logFatal
	defer func() { logFatal = origLogFatal }()
	
	called := false
	logFatal = func(format string, v ...interface{}) {
		called = true
	}

	main()

	if !called {
		t.Error("Expected logFatal to be called, but it was not")
	}
}
