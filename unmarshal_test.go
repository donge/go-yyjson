package main

import (
	"encoding/json"
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
