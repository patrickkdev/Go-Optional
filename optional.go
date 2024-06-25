package optional

import "errors"

// Optional represents an optional value.
type Optional[T any] struct {
	value *T
}

// New creates a new Optional containing the given value.
func New[T any](value T) Optional[T] {
	return Optional[T]{value: &value}
}

// Empty creates a new empty Optional.
func Empty[T any]() Optional[T] {
	return Optional[T]{value: nil}
}

// IsPresent returns true if the Optional contains a value.
func (opt Optional[T]) IsPresent() bool {
	return opt.value != nil
}

// Get returns the value if present, or an error if not.
func (opt Optional[T]) Get() (T, error) {
	if opt.value != nil {
		return *opt.value, nil
	}
	var zero T
	return zero, errors.New("no value present")
}

// OrElse returns the value if present, or the provided default value if not.
func (opt Optional[T]) OrElse(defaultValue T) T {
	if opt.value != nil {
		return *opt.value
	}
	return defaultValue
}
