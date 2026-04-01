package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMArray(t *testing.T) {
	t.Run("BasicOperations", func(t *testing.T) {
		arr := NewMArray()
		assert.NotNil(t, arr)
		assert.Equal(t, 0, arr.size)

		// Test MADDEND
		arr.MADDEND("first")
		assert.Equal(t, 1, arr.size)

		arr.MADDEND("second")
		assert.Equal(t, 2, arr.size)

		// Test MGETINDEX valid
		val, err := arr.MGETINDEX(0)
		require.NoError(t, err)
		assert.Equal(t, "first", val)

		val, err = arr.MGETINDEX(1)
		require.NoError(t, err)
		assert.Equal(t, "second", val)

		// Test MLENGTH
		assert.Equal(t, 2, arr.MLENGTH())
	})

	t.Run("MADDINDEX", func(t *testing.T) {
		arr := NewMArray()
		arr.MADDEND("A")
		arr.MADDEND("C")

		arr.MADDINDEX(1, "B")
		assert.Equal(t, 3, arr.size)

		val, err := arr.MGETINDEX(1)
		require.NoError(t, err)
		assert.Equal(t, "B", val)
	})

	t.Run("ErrorCases", func(t *testing.T) {
		arr := NewMArray()
		arr.MADDEND("A")

		// Capture stdout for error messages
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// MADDINDEX invalid index
		arr.MADDINDEX(5, "B")

		// MREMOVEINDEX invalid indices
		arr.MREMOVEINDEX(5)
		arr.MREMOVEINDEX(-1)

		// MREPLACEINDEX invalid indices
		arr.MREPLACEINDEX(5, "new")
		arr.MREPLACEINDEX(-1, "new")

		// MGETINDEX invalid
		_, err := arr.MGETINDEX(5)
		assert.Error(t, err)
		_, err = arr.MGETINDEX(-1)
		assert.Error(t, err)

		w.Close()
		_, _ = io.ReadAll(r)
		os.Stdout = oldStdout

		// Original state preserved
		assert.Equal(t, 1, arr.size)
		val, err := arr.MGETINDEX(0)
		require.NoError(t, err)
		assert.Equal(t, "A", val)
	})

	t.Run("MREPLACEINDEX", func(t *testing.T) {
		arr := NewMArray()
		arr.MADDEND("old")

		// Capture stdout
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		arr.MREPLACEINDEX(0, "new")

		w.Close()
		_, _ = io.ReadAll(r)
		os.Stdout = oldStdout

		val, err := arr.MGETINDEX(0)
		require.NoError(t, err)
		assert.Equal(t, "new", val)
	})

	t.Run("MCLEAR", func(t *testing.T) {
		arr := NewMArray()
		arr.MADDEND("A")
		arr.MADDEND("B")

		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		arr.MCLEAR()

		w.Close()
		_, _ = io.ReadAll(r)
		os.Stdout = oldStdout

		assert.Equal(t, 0, arr.size)
		assert.Equal(t, 0, len(arr.data))
	})

	t.Run("MPRINT", func(t *testing.T) {
		arr := NewMArray()

		// Test empty print
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		arr.MPRINT()

		w.Close()
		_, _ = io.ReadAll(r)
		os.Stdout = oldStdout

		// Test non-empty print
		arr.MADDEND("Hello")
		arr.MADDEND("World")

		oldStdout = os.Stdout
		r, w, _ = os.Pipe()
		os.Stdout = w

		arr.MPRINT()

		w.Close()
		_, _ = io.ReadAll(r)
		os.Stdout = oldStdout
	})

	t.Run("SetGetAll", func(t *testing.T) {
		arr := NewMArray()
		arr.MADDEND("A")
		arr.MADDEND("B")

		data := arr.GetAll()
		assert.Equal(t, []string{"A", "B"}, data)

		newArr := NewMArray()
		newArr.SetAll([]string{"X", "Y"})
		assert.Equal(t, 2, newArr.size)
		assert.Equal(t, []string{"X", "Y"}, newArr.GetAll())
	})

	t.Run("MREMOVEINDEX", func(t *testing.T) {
		arr := NewMArray()
		arr.MADDEND("A")
		arr.MADDEND("B")
		arr.MADDEND("C")

		arr.MREMOVEINDEX(1) // Remove "B"
		assert.Equal(t, 2, arr.size)

		val, err := arr.MGETINDEX(0)
		require.NoError(t, err)
		assert.Equal(t, "A", val)

		val, err = arr.MGETINDEX(1)
		require.NoError(t, err)
		assert.Equal(t, "C", val)
	})
}
