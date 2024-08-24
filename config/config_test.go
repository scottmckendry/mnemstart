package config

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_ENV", "test_value")

	result := getEnv("TEST_ENV", "fallback_value")
	if result != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", result)
	}

	result = getEnv("NON_EXISTENT_ENV", "fallback_value")
	if result != "fallback_value" {
		t.Errorf("Expected 'fallback_value', got '%s'", result)
	}

	t.Cleanup(func() {
		os.Unsetenv("TEST_ENV")
	})
}

func TestGetEnvAsInt(t *testing.T) {
	os.Setenv("TEST_ENV", "123")

	result := getEnvAsInt("TEST_ENV", 456)
	if result != 123 {
		t.Errorf("Expected 123, got %d", result)
	}

	result = getEnvAsInt("NON_EXISTENT_ENV", 456)
	if result != 456 {
		t.Errorf("Expected 456, got %d", result)
	}

	t.Cleanup(func() {
		os.Unsetenv("TEST_ENV")
	})
}

func TestGetEnvAsBool(t *testing.T) {
	os.Setenv("TEST_ENV", "true")

	result := getEnvAsBool("TEST_ENV", false)
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}

	result = getEnvAsBool("NON_EXISTENT_ENV", false)
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}

	t.Cleanup(func() {
		os.Unsetenv("TEST_ENV")
	})
}
