package main // Benchmark test for the Fibonacci function

import (
	"encoding/base64"
	"testing"
)

func BenchmarkBase64Encoding(b *testing.B) {
	data := []byte(`{"id":15,"username":"admin","layers":["GRAPHRASTER:eu_composite_flood_risk"],"resultsPerPage":2,"page":0,"operation":"enhance"}`) // Replace with your data or generate dynamically
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base64.StdEncoding.EncodeToString(data)
	}
}

// Run this benchmark with `go test -bench=.` command in the terminal.
