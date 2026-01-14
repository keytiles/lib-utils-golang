package kt_utils_test

import (
	"testing"

	"github.com/keytiles/lib-utils-golang/pkg/kt_utils"
	"github.com/stretchr/testify/assert"
)

type Person struct {
}

func TestStringSimpleResolve_fullVars(t *testing.T) {
	// ---- GIVEN
	tpl := "My name is {name} I am {age} years old. Am I a man? Well, this is {isMan} :-) I live in {place} - I am a(n) {profession} at {workplace}"
	vars := map[string]any{"name": "Attila", "age": 50, "isMan": true, "place": "Munich", "profession": "Software Architect", "workplace": "Keytiles"}

	// ---- WHEN
	resolved := kt_utils.StringSimpleResolve(tpl, vars)

	// ---- THEN
	assert.Equal(
		t,
		"My name is Attila I am 50 years old. Am I a man? Well, this is true :-) I live in Munich - I am a(n) Software Architect at Keytiles",
		resolved,
	)
}

func TestStringSimpleResolve_partialVars(t *testing.T) {
	// ---- GIVEN
	tpl := "My name is {name} I am {age} years old. Am I a man? Well, this is {isMan} :-) I live in {place} - I am a(n) {profession} at {workplace}"
	vars := map[string]any{"name": "Attila", "age": 50, "isMan": true, "workplace": "Keytiles"}

	// ---- WHEN
	resolved := kt_utils.StringSimpleResolve(tpl, vars)

	// ---- THEN
	assert.Equal(
		t,
		"My name is Attila I am 50 years old. Am I a man? Well, this is true :-) I live in {place} - I am a(n) {profession} at Keytiles",
		resolved,
	)
}

func TestStringSimpleResolve_notInitializedVars(t *testing.T) {
	// ---- GIVEN
	tpl := "My name is {name} I am {age} years old. Am I a man? Well, this is {isMan} :-) I live in {place} - I am a(n) {profession} at {workplace}"
	var vars map[string]any

	// ---- WHEN
	resolved := kt_utils.StringSimpleResolve(tpl, vars)

	// ---- THEN
	assert.Equal(t, tpl, resolved)
}
