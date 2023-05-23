package xhttp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewHttp(t *testing.T) {
	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	h := NewHttp(WithTimeout(time.Second), WithHeaders(headers))

	assert.Equal(t, time.Second, h.Timeout)
	assert.Equal(t, headers, h.Headers)
}
