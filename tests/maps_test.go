package kt_utils_test

import (
	"testing"

	"github.com/keytiles/lib-utils-golang/v2/pkg/kt_utils"
	"github.com/stretchr/testify/assert"
)

func Test_MapKeys(t *testing.T) {
	// ---- GIVEN
	m := map[string]int{"a": 1, "b": 6, "c": 6, "d": 1, "e": 3}
	// ---- WHEN
	keys := kt_utils.MapKeys(m)
	// ---- THEN
	assert.Equal(t, 5, len(keys))
	assert.Contains(t, keys, "a")
	assert.Contains(t, keys, "b")
	assert.Contains(t, keys, "c")
	assert.Contains(t, keys, "d")
	assert.Contains(t, keys, "e")
}

func Test_MapKeys_emptymap(t *testing.T) {
	// ---- GIVEN
	m := map[string]int{}
	// ---- WHEN
	keys := kt_utils.MapKeys(m)
	// ---- THEN
	assert.Equal(t, 0, len(keys))
}

func Test_MapValuesSet(t *testing.T) {
	// ---- GIVEN
	m := map[string]int{"a": 1, "b": 6, "c": 6, "d": 1, "e": 3}
	// ---- WHEN
	values := kt_utils.MapValuesSet(m)
	// ---- THEN
	assert.Equal(t, 3, values.Size())
	assert.True(t, values.Contains(1))
	assert.True(t, values.Contains(3))
	assert.True(t, values.Contains(6))
}

func Test_MapValuesSet_emptymap(t *testing.T) {
	// ---- GIVEN
	m := map[string]int{}
	// ---- WHEN
	values := kt_utils.MapValuesSet(m)
	// ---- THEN
	assert.Equal(t, 0, values.Size())
}
