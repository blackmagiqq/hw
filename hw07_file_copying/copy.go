package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFileNotFound          = errors.New("file not found")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	from, err := os.Open(fromPath)
	if err != nil {
		return ErrFileNotFound
	}
	defer from.Close()
	fromStat, _ := from.Stat()
	if fromStat.Size() == 0 {
		return ErrUnsupportedFile
	}
	if offset > fromStat.Size() {
		return errors.New("offset cannot be greater than file size")
	}
	if limit > fromStat.Size() {
		return errors.New("limit cannot be greater than file size")
	}
	if limit == 0 {
		limit = fromStat.Size() - offset
	}
	to, err := os.OpenFile(toPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer to.Close()
	from.Seek(offset, 0)
	progressReader := &ProgressReader{
		Reader: io.LimitReader(from, limit),
		total:  limit,
	}
	_, err = to.ReadFrom(progressReader)
	if err != nil {
		return err
	}
	return nil
}
