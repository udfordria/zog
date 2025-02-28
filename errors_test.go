package zog

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	p "github.com/udfordria/zog/internals"
)

func TestZogIssueString(t *testing.T) {
	var err p.ZogIssue = &p.ZogErr{
		C:       "test",
		ParamsM: map[string]any{},
		Typ:     "string",
		Val:     "asda",
		Msg:     "asda",
	}

	assert.Equal(t, err.String(), err.Error())
}

func TestZogIssueUnwrap(t *testing.T) {
	var err p.ZogIssue = &p.ZogErr{
		Err: errors.New("test"),
	}

	assert.Equal(t, err.Unwrap().Error(), "test")
}
