package gojson

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_upperCamelToUnderScore(t *testing.T) {
	require.Equal(t, upperCamelToUnderScore("XxYy"), "xx_yy")
	require.Equal(t, upperCamelToUnderScore("XxYY"), "xx_y_y")
	require.Equal(t, upperCamelToUnderScore("XxY_Y"), "xx_y_y")
	require.Equal(t, upperCamelToUnderScore("xxYY"), "xx_y_y")
	require.Equal(t, upperCamelToUnderScore("Xx_YY"), "xx_y_y")
	require.Equal(t, upperCamelToUnderScore("xx_yy"), "xx_yy")
}

func Test_underScoreToUpperCamel(t *testing.T) {
	require.Equal(t, underScoreToUpperCamel("xx_yy"), "XxYy")
	require.Equal(t, underScoreToUpperCamel("xx_y_y"), "XxYY")
	require.Equal(t, underScoreToUpperCamel("Xx_Y_y"), "Xx_YY")
}
