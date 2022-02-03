package protovalidator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayIsUnique(t *testing.T) {
	t.Run("positive int32", func(t *testing.T) {
		a := []int32{1, 2, 3, 4}
		ok := SliceIsUniqueInt32(a)
		require.True(t, ok)
	})
	t.Run("positive int64", func(t *testing.T) {
		a := []int64{1, 2, 3, 4}
		ok := SliceIsUniqueInt64(a)
		require.True(t, ok)
	})
	t.Run("positive uint32", func(t *testing.T) {
		a := []uint32{1, 2, 3, 4}
		ok := SliceIsUniqueUint32(a)
		require.True(t, ok)
	})
	t.Run("positive uint64", func(t *testing.T) {
		a := []uint64{1, 2, 3, 4}
		ok := SliceIsUniqueUint64(a)
		require.True(t, ok)
	})
	t.Run("positive string", func(t *testing.T) {
		a := []string{"s1", "s2", "s3", "s4"}
		ok := SliceIsUniqueString(a)
		require.True(t, ok)
	})
	t.Run("positive bool", func(t *testing.T) {
		a := []bool{true, false}
		ok := SliceIsUniqueBool(a)
		require.True(t, ok)
	})
	t.Run("positive float64", func(t *testing.T) {
		a := []float64{1.1, 2.2, 3.3, 4.4}
		ok := SliceIsUniqueFloat64(a)
		require.True(t, ok)
	})
	t.Run("positive float32", func(t *testing.T) {
		a := []float32{1.1, 2.2, 3.3, 4.4}
		ok := SliceIsUniqueFloat32(a)
		require.True(t, ok)
	})

	t.Run("negative int32", func(t *testing.T) {
		a := []int32{1, 2, 1, 4}
		ok := SliceIsUniqueInt32(a)
		require.False(t, ok)
	})
	t.Run("negative int64", func(t *testing.T) {
		a := []int64{1, 2, 1, 4}
		ok := SliceIsUniqueInt64(a)
		require.False(t, ok)
	})
	t.Run("negative uint32", func(t *testing.T) {
		a := []uint32{1, 2, 1, 4}
		ok := SliceIsUniqueUint32(a)
		require.False(t, ok)
	})
	t.Run("negative uint64", func(t *testing.T) {
		a := []uint64{1, 2, 1, 4}
		ok := SliceIsUniqueUint64(a)
		require.False(t, ok)
	})
	t.Run("negative string", func(t *testing.T) {
		a := []string{"s1", "s2", "s1", "s4"}
		ok := SliceIsUniqueString(a)
		require.False(t, ok)
	})
	t.Run("negative bool", func(t *testing.T) {
		a := []bool{true, true}
		ok := SliceIsUniqueBool(a)
		require.False(t, ok)
	})
	t.Run("negative float64", func(t *testing.T) {
		a := []float64{1.1, 2.2, 1.1, 4.4}
		ok := SliceIsUniqueFloat64(a)
		require.False(t, ok)
	})
	t.Run("negative float32", func(t *testing.T) {
		a := []float32{1.1, 2.2, 1.1, 4.4}
		ok := SliceIsUniqueFloat32(a)
		require.False(t, ok)
	})
}
