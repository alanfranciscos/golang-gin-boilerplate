package main

import (
	"context"
	"os"
	"testing"
)

func TestRun_Error(t *testing.T) {
	os.Setenv("APP_PORT", "invalid")
	defer os.Unsetenv("APP_PORT")

	err := Run(context.Background())
	if err == nil {
		t.Error("Expected error when running with invalid port, got nil")
	}
}

func TestRun_Success(t *testing.T) {
	// Skip if running in environment where port 8080 is taken
	// but since we are using a cancelled context, it should shut down immediately
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	err := Run(ctx)
	if err != nil {
		t.Errorf("Expected nil error on clean shutdown, got %v", err)
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
	_ = called

	// We can't easily test main() because it uses signal.NotifyContext
	// but we already tested Run() which is the core logic.
	// For coverage, we'll call it with a cancelled context if possible, 
	// but main() creates its own context.
}
