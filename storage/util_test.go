package storage

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
	"time"
)

func TestMarshallingDuration_UnmarshalText(t *testing.T) {
	type Container struct {
		Value MarshallingDuration
	}

	expected := Container{Value: MarshallingDuration(time.Hour * 3)}

	t.Run("roundtrip", func(t *testing.T) {
		// Marshal
		data, _ := yaml.Marshal(expected)
		assert.Equal(t, string(data), "value: 3h0m0s\n")

		// Unmarshal
		var actual Container
		err := yaml.Unmarshal(data, &actual)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("minimal", func(t *testing.T) {
		var actual Container
		err := yaml.Unmarshal([]byte("value: 3h"), &actual)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("empty", func(t *testing.T) {
		var actual Container
		err := yaml.Unmarshal([]byte("value: "), &actual)
		assert.NoError(t, err)
		assert.Equal(t, time.Duration(0), time.Duration(actual.Value))
	})
}
