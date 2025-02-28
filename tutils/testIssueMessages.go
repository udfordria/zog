package tutils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/udfordria/zog/i18n/en"
	"github.com/udfordria/zog/internals"
	"github.com/udfordria/zog/zconst"
)

func VerifyDefaultIssueMessages(t *testing.T, errs internals.ZogIssueList) {
	for _, err := range errs {
		c := err.Code()
		m, ok := en.Map[err.Dtype()][c]
		if !ok {
			m, ok = en.Map[err.Dtype()][zconst.IssueCodeFallback]
			if !ok {
				panic(fmt.Sprintf("no fallback message for type %s", err.Dtype()))
			}
		}
		prefix := strings.Split(m, "{{")[0]
		postfix := prefix
		if strings.Contains(m, "}}") {
			postfix = strings.Split(m, "}}")[1]
		}
		assert.True(t, strings.HasPrefix(err.Message(), prefix))
		assert.True(t, strings.HasSuffix(err.Message(), postfix))
	}
}

func VerifyDefaultIssueMessagesMap(t *testing.T, errs internals.ZogIssueMap) {
	for _, errList := range errs {
		VerifyDefaultIssueMessages(t, errList)
	}
}
