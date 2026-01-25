package main

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	segmentjson "github.com/segmentio/encoding/json"
	"github.com/tidwall/gjson"
	"testing"
)

func BenchmarkComprehensiveUnmarshal(b *testing.B) {
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

func BenchmarkComprehensiveStdUnmarshal(b *testing.B) {
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

func BenchmarkComprehensiveUnmarshalOptimized(b *testing.B) {
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

func BenchmarkComprehensiveSonicUnmarshal(b *testing.B) {
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

func BenchmarkComprehensiveGjsonParse(b *testing.B) {
	// 准备测试数据
	jsonData := `{"name":"Mash","star":4,"hits":[1,2,3,4], "a": { "b": {"c":"d"}}}`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 使用 gjson 解析并获取所有主要值来模拟完整的unmarshal
		result := gjson.Parse(jsonData)

		// 获取所有值来模拟完整的map[string]interface{}转换
		name := result.Get("name").String()
		star := result.Get("star").Int()
		hits := result.Get("hits").Array()
		a_val := result.Get("a.b.c").String()

		// 验证解析是否成功
		if !result.Exists() || name == "" || star == 0 || len(hits) == 0 || a_val == "" {
			b.Fatalf("gjson.Parse or value extraction failed")
		}
	}
}

func BenchmarkComprehensiveJsoniterUnmarshal(b *testing.B) {
	// 准备测试数据
	jsonData := []byte(`{"name":"Mash","star":4,"hits":[1,2,3,4], "a": { "b": {"c":"d"}}}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 使用 json-iterator 的 Unmarshal 函数
		var data map[string]interface{}
		if err := jsoniter.Unmarshal(jsonData, &data); err != nil {
			b.Fatalf("jsoniter.Unmarshal error: %v", err)
		}
	}
}

func BenchmarkComprehensiveGojsonUnmarshal(b *testing.B) {
	// 准备测试数据
	jsonData := []byte(`{"name":"Mash","star":4,"hits":[1,2,3,4], "a": { "b": {"c":"d"}}}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 使用 go-json 的 Unmarshal 函数
		var data map[string]interface{}
		if err := gojson.Unmarshal(jsonData, &data); err != nil {
			b.Fatalf("gojson.Unmarshal error: %v", err)
		}
	}
}

func BenchmarkComprehensiveSegmentjsonUnmarshal(b *testing.B) {
	// 准备测试数据
	jsonData := []byte(`{"name":"Mash","star":4,"hits":[1,2,3,4], "a": { "b": {"c":"d"}}}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 使用 segmentio json 的 Unmarshal 函数
		var data map[string]interface{}
		if err := segmentjson.Unmarshal(jsonData, &data); err != nil {
			b.Fatalf("segmentjson.Unmarshal error: %v", err)
		}
	}
}
