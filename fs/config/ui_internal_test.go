package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckPassword(t *testing.T) {
	// Empty password should give an error
	_, err := checkPassword("")
	require.Error(t, err)

	// Whitespace only should give an error
	_, err = checkPassword("  \t  ")
	require.Error(t, err)

	// Invalid utf8 should give an error
	_, err = checkPassword(string([]byte{0xff, 0xfe, 0xfd}) + "abc")
	require.Error(t, err)

	// Passwords must be returned verbatim, in particular without
	// Unicode normalization, so that obscured backend passwords reveal
	// to exactly what the user typed - see #9507. NFKC normalization
	// would turn ª (U+00AA) into a.
	pw := "ËQÖ4];giª±;<H'õPVzú6j¢ÔBí}×_-D%}"
	got, err := checkPassword(pw)
	require.NoError(t, err)
	assert.Equal(t, pw, got)
}
