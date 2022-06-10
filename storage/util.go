package storage

import (
	"fmt"
	"time"
)

// MarshallingDuration wraps time.Duration so it can be unmarshalled from YAML
type MarshallingDuration time.Duration

func (d MarshallingDuration) MarshalText() ([]byte, error) {
	return []byte(time.Duration(d).String()), nil
}

// UnmarshalText parses a string representation of time.Duration into time.Duration
func (d *MarshallingDuration) UnmarshalText(text []byte) error {
	value, err := time.ParseDuration(string(text))
	*d = MarshallingDuration(value)
	if err != nil {
		return fmt.Errorf("invalid duration '%s': %w", string(text), err)
	}
	return nil
}
