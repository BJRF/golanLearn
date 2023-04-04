package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const (
	str = "efwaefnurgnrehgepbnrebewnbgblasjfnowbgwooihfunw"
	cnt = 10000
)

// BenchmarkPlusConcat + 拼接
func BenchmarkPlusConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ss := ""
		for i := 0; i < cnt; i++ {
			ss += str
		}
	}
}

// BenchmarkSprintfConcat sprintf拼接
func BenchmarkSprintfConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ss := ""
		for i := 0; i < cnt; i++ {
			ss = fmt.Sprintf("%s%s", ss, str)
		}
	}
}

// BenchmarkBuilderConcat stringbuilder 拼接
func BenchmarkBuilderConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for i := 0; i < cnt; i++ {
			builder.WriteString(str)
		}
		builder.String()
	}
}

// BenchmarkBufferConcat stringbuilder 拼接
func BenchmarkBufferConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		for i := 0; i < cnt; i++ {
			buf.WriteString(str)
		}
		buf.String()
	}
}

// BenchmarkAppendConcat append 拼接
func BenchmarkAppendConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 0)
		for i := 0; i < cnt; i++ {
			buf = append(buf, str...)
		}
	}
}
