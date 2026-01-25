package main

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	"testing"
)

func BenchmarkUnmarshal(b *testing.B) {
	// 准备测试数据
	jsonData := []byte(`{"name":"Mash","star":4,"hits":[1,2,3,4], "a": { "b": {"c":"d"}}}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 使用 Unmarshal 函数
		var data map[string]interface{}
		if err := Unmarshal(jsonData, &data); err != nil {
			b.Fatalf("Unmarshal error: %v", err)
		}
	}
}

func BenchmarkStdUnmarshal(b *testing.B) {
	// 准备测试数据
	jsonData := []byte(`{"name":"Mash","star":4,"hits":[1,2,3,4], "a": { "b": {"c":"d"}}}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 使用标准库的 json.Unmarshal 函数
		var data map[string]interface{}
		if err := json.Unmarshal(jsonData, &data); err != nil {
			b.Fatalf("json.Unmarshal error: %v", err)
		}
	}
}

func BenchmarkUnmarshalOptimized(b *testing.B) {
	// 准备测试数据
	jsonData := []byte(`{"name":"Mash","star":4,"hits":[1,2,3,4], "a": { "b": {"c":"d"}}}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 使用优化版本的 UnmarshalOptimized 函数
		var data map[string]interface{}
		if err := UnmarshalOptimized(jsonData, &data); err != nil {
			b.Fatalf("UnmarshalOptimized error: %v", err)
		}
	}
}

func BenchmarkSonicUnmarshal(b *testing.B) {
	// 准备测试数据
	jsonData := []byte(`{"name":"Mash","star":4,"hits":[1,2,3,4], "a": { "b": {"c":"d"}}}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 使用 sonic 的 Unmarshal 函数
		var data map[string]interface{}
		if err := sonic.Unmarshal(jsonData, &data); err != nil {
			b.Fatalf("sonic.Unmarshal error: %v", err)
		}
	}
}
