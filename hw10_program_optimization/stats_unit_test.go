package hw10programoptimization

import (
	"bytes"
	"testing"
)

func BenchmarkGetDomainStat(b *testing.B) {
	b.ReportAllocs()
	data := []byte(`{"Email": "user1@example.com"}
{"Email": "user2@example.com"}
{"Email": "user3@test.com"}
{"Email": "user4@example.com"}
{"Email": "user5@test.com"}`)

	domain := "example.com"

	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(data)
		_, err := GetDomainStat(r, domain)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
