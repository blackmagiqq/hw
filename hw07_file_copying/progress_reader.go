package main

import (
	"fmt"
	"io"
)

type ProgressReader struct {
	io.Reader       // Встроенный интерфейс Reader
	total     int64 // Общий размер файла
	current   int64 // Текущее количество прочитанных байт
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.Reader.Read(p)
	pr.current += int64(n)

	percent := float64(pr.current) / float64(pr.total) * 100
	fmt.Printf("\rКопирование: %.2f%%", percent)

	return n, err
}
