package course

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRawCBRCourse(t *testing.T) {
	req := getRawCBRCourse()

	assert.IsType(t, req, CBRRequest{})
	assert.Equal(t, len(req.Data), 34)
}

func TestGetCBRCourse(t *testing.T) {
	req := GetCBRCourse()

	assert.IsType(t, req, CBRData{})
	assert.Equal(t, len(req.Data), 34)
}
