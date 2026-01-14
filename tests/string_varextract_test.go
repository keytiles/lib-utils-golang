package kt_utils_test

import (
	"testing"

	"github.com/keytiles/lib-sets-golang/ktsets"
	"github.com/keytiles/lib-utils-golang/pkg/kt_utils"
	"github.com/stretchr/testify/assert"
)

func TestExtractVariableNamesFromString(t *testing.T) {
	// ---- GIVEN
	str := "This is a text with {var1} and {var2} variables. And we repeat {var1} again"
	// ---- WHEN
	vars := kt_utils.StringExtractVariableNames(str)
	// ---- THEN
	assert.True(t, vars.Equals(ktsets.NewSet("var1", "var2")))
}
