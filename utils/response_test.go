package utils

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestResp(t *testing.T) {
	var resp Response
	res := resp.Success("data").WithDesc("success desc")
	assert.Equal(t, res.Desc, "success desc")
	res = resp.Failure().WithDesc("failure desc")
	assert.Equal(t, res.Desc, "failure desc")
}
