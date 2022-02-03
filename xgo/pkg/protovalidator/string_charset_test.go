package protovalidator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringIsAlpha(t *testing.T) {
	for x := 'a'; x < 'z'; x++ {
		ok := StringIsAlpha(string(x))
		require.True(t, ok, string(x))
	}
	for x := 'A'; x < 'Z'; x++ {
		ok := StringIsAlpha(string(x))
		require.True(t, ok, string(x))
	}

	for x := byte(0); x < 'A'; x++ {
		ok := StringIsAlpha(string(x))
		require.False(t, ok, string(x))
	}

	for x := byte(91); x < 'a'; x++ {
		ok := StringIsAlpha(string(x))
		require.False(t, ok, string(x))
	}

	for x := byte(123); x < 255; x++ {
		ok := StringIsAlpha(string(x))
		require.False(t, ok, string(x))
	}
}

func TestStringIsNumber(t *testing.T) {
	for x := '0'; x < '9'; x++ {
		ok := StringIsNumber(string(x))
		require.True(t, ok, string(x))
	}

	for x := byte(0); x < '0'; x++ {
		ok := StringIsNumber(string(x))
		require.False(t, ok, string(x))
	}

	for x := byte(58); x < 255; x++ {
		ok := StringIsNumber(string(x))
		require.False(t, ok, string(x))
	}
}

func TestStringIsAscii(t *testing.T) {
	for x := 'a'; x < 'z'; x++ {
		ok := StringIsAlphaNumber(string(x))
		require.True(t, ok, string(x))
	}
	for x := 'A'; x < 'Z'; x++ {
		ok := StringIsAlphaNumber(string(x))
		require.True(t, ok, string(x))
	}
	for x := '0'; x < '9'; x++ {
		ok := StringIsAlphaNumber(string(x))
		require.True(t, ok, string(x))
	}

	for x := byte(0); x < '0'; x++ {
		ok := StringIsAlphaNumber(string(x))
		require.False(t, ok, string(x))
	}
	for x := byte(58); x < 'A'; x++ {
		ok := StringIsAlphaNumber(string(x))
		require.False(t, ok, string(x))
	}
	for x := byte(91); x < 'a'; x++ {
		ok := StringIsAlphaNumber(string(x))
		require.False(t, ok, string(x))
	}
	for x := byte(123); x < 255; x++ {
		ok := StringIsAlphaNumber(string(x))
		require.False(t, ok, string(x))
	}
}
