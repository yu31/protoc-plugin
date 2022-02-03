package jsonencoder

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	enc := New(0)
	require.NotNil(t, enc)
	//require.NotNil(t, enc.buf)
}

func TestEncoder_Bytes(t *testing.T) {
	enc := New(0)
	enc.AppendObjectBegin()
	//enc.AddString("s1", "sv1")
	//enc.AddInt64("int64", 64)
	enc.AppendObjectEnd()

	b := enc.Bytes()
	require.NotNil(t, b)
	println(string(b))

	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	require.Nil(t, err)
	//require.Equal(t, m["s1"], "sv1")
	//require.Equal(t, m["int64"], float64(64))

}
