package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	p "github.com/udfordria/zog/internals"
	zconst "github.com/udfordria/zog/zconst"
)

func TestDefaultIssueFormatter(t *testing.T) {
	tests := []struct {
		input p.ZogIssue
		want  string
	}{
		{input: &p.ZogErr{C: zconst.IssueCodeRequired, Typ: zconst.TypeString}, want: DefaultErrMsgMap[zconst.TypeString][zconst.IssueCodeRequired]},
		{input: &p.ZogErr{C: zconst.IssueCodeRequired, Typ: zconst.TypeString, Msg: "DON'T OVERRIDE ME"}, want: "DON'T OVERRIDE ME"},
		{input: &p.ZogErr{C: "INVALID_ERR_CODE", Typ: zconst.TypeString}, want: "string is invalid"},
	}

	for _, test := range tests {
		IssueFormatter(test.input, nil)
		assert.Equal(t, test.want, test.input.Message())
	}
}
