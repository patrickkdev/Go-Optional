package optional

import (
	"testing"
)

func TestNew(t *testing.T) {
	opt := New(42)
	if !opt.IsPresent() {
		t.Error("Expected value to be present")
	}

	value, err := opt.Get()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != 42 {
		t.Errorf("Expected 42, got %v", value)
	}
}

func TestEmpty(t *testing.T) {
	opt := Empty[int]()
	if opt.IsPresent() {
		t.Error("Expected no value to be present")
	}

	_, err := opt.Get()
	if err == nil {
		t.Error("Expected an error")
	}
}

func TestOrElse(t *testing.T) {
	opt := Empty[int]()
	defaultValue := opt.OrElse(100)
	if defaultValue != 100 {
		t.Errorf("Expected 100, got %v", defaultValue)
	}

	opt = New(42)
	value := opt.OrElse(100)
	if value != 42 {
		t.Errorf("Expected 42, got %v", value)
	}
}

func TestOptionalString(t *testing.T) {
	opt := New("Hello, World!")
	if !opt.IsPresent() {
		t.Error("Expected value to be present")
	}

	value, err := opt.Get()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if value != "Hello, World!" {
		t.Errorf("Expected 'Hello, World!', got %v", value)
	}
}
