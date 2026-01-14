package kt_utils

import (
	"fmt"
	"regexp"

	"github.com/keytiles/lib-sets-golang/ktsets"
)

var (
	VARIABLE_MATCHER = regexp.MustCompile(`\{([a-zA-Z0-9_]+)\}`)
)

// We have a string with variable names in it (Python style), e.g. "My string with {var1} and {var2} variables."
// This function can resolve this string with the provided map of values. If there is no entry for a specific variable
// in the map then the variable stays in place unresolved.
func StringSimpleResolve(input string, vars map[string]any) string {
	return VARIABLE_MATCHER.ReplaceAllStringFunc(input, func(match string) string {
		key := VARIABLE_MATCHER.FindStringSubmatch(match)[1]
		if val, ok := vars[key]; ok {
			return fmt.Sprint(val)
		}
		return match // leave unchanged if missing
	})
}

// We have a string with variable names in it (Python style), e.g. "My string with {var1} and {var2} variables."
// This method returns all variable names it finds in the text, in our case {"var1", "var2"} would be returned
func StringExtractVariableNames(s string) ktsets.Set[string] {
	matches := VARIABLE_MATCHER.FindAllStringSubmatch(s, -1)
	result := ktsets.NewSetWithCapacity[string](len(matches))
	for _, match := range matches {
		if len(match) > 1 {
			result.Add(match[1]) // match[1] is the first group
		}
	}
	return result
}
