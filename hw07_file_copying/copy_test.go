package main

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)


const (
	fromFileName = "testfile"
	fromContent  = "Hello, Otus!"
)

func TestCopy(t *testing.T) {
	t.Run("negative - file from not found", func(t *testing.T) {
		err := Copy("foo", "bar", 0, 0)
		require.ErrorAs(t, err, &ErrUnsupportedFile)
	})
	t.Run("negative - offset is greater than file", func(t *testing.T) {
		fromContentLength := len(fromContent)

		from, _ := os.Create(fromFileName)
		defer from.Close()
		defer func() {
			os.Remove(fromFileName)
		}()
		from.WriteString(fromContent)

		pwd, _ := os.Getwd()

		err := Copy(path.Join(pwd, fromFileName), "bar", int64(fromContentLength)+1, 0)
		require.ErrorAs(t, err, &ErrOffsetLength)
	})
	t.Run("negative - limit is greater than file", func(t *testing.T) {
		fromContentLength := len(fromContent)

		from, _ := os.Create(fromFileName)
		defer from.Close()
		defer func() {
			os.Remove(fromFileName)
		}()
		from.WriteString(fromContent)

		pwd, _ := os.Getwd()
		err := Copy(path.Join(pwd, fromFileName), "bar", 0, int64(fromContentLength)+1)
		require.ErrorAs(t, err, &ErrLimitLength)
	})
	t.Run("positive - simple copy", func(t *testing.T) {
		toFileName := "testfileto"

		from, _ := os.Create(fromFileName)
		defer from.Close()
		defer func() {
			os.Remove(fromFileName)
		}()
		from.WriteString(fromContent)

		pwd, _ := os.Getwd()
		err := Copy(path.Join(pwd, fromFileName), path.Join(pwd, toFileName), 0, 0)
		if err != nil {
			t.Fatalf("copy fail: %v", err)
		}

		to, err := os.OpenFile(toFileName, os.O_RDONLY, 0o222)
		if err != nil {
			t.Fatalf("file not created: %v", err)
		}
		defer to.Close()
		defer func() {
			os.Remove(toFileName)
		}()
		actualContent := make([]byte, len(fromContent))
		to.Read(actualContent)
		require.Equal(t, fromContent, string(actualContent))
	})
}
