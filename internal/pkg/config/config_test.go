package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_UpdateFromBytes(t *testing.T) {
	// CAUTION: YAML here should be indented with spaces, tabs may cause exceptions.
	cfg := `
debug: false
token: "mocked-token"
`
	c := &Config{}
	err := c.UpdateFromBytes([]byte(cfg))

	if nil != err {
		t.Error(err)
	}

	assert.False(t, c.DebugMode)
	assert.Equal(t, "mocked-token", c.Token)
}
