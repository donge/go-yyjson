package main

import (
	"github.com/valyala/fastjson"
	"testing"
)

func BenchmarkComprehensiveFastjsonParse(b *testing.B) {
	// 准备测试数据
	jsonData := string(`{"name":"Mash","star":4,"hits":[1,2,3,4], "a": { "b": {"c":"d"}}}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 使用 fastjson 解析
		parsed, err := fastjson.Parse(jsonData)
		if err != nil {
			b.Fatalf("fastjson.Parse error: %v", err)
		}

		// 获取值来模拟实际使用场景
		name := parsed.GetStringBytes("name")
		star := parsed.GetInt("star")
		hits := parsed.GetArray("hits")
		a_val := parsed.GetStringBytes("a", "b", "c")

		// 验证解析是否成功
		if len(name) == 0 || star == 0 || hits == nil || len(a_val) == 0 {
			b.Fatalf("fastjson.Parse or value extraction failed")
		}
	}
}

func BenchmarkLargeFastjsonParse(b *testing.B) {
	jsonStr := string(largeJSONData)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 使用 fastjson 解析大JSON
		parsed, err := fastjson.Parse(jsonStr)
		if err != nil {
			b.Fatalf("fastjson.Parse error: %v", err)
		}

		// 获取主要的嵌套值来模拟实际使用场景
		userName := parsed.GetStringBytes("user", "name")
		totalViews := parsed.GetInt("analytics", "total_views")
		postCount := parsed.GetInt("analytics", "total_posts")
		topPostViews := parsed.GetInt("analytics", "top_performing_post", "views")
		unreadCount := parsed.GetInt("notifications", "unread_count")

		// 验证解析是否成功
		if len(userName) == 0 || totalViews == 0 || postCount == 0 || topPostViews == 0 || unreadCount == 0 {
			b.Fatalf("fastjson.Parse or value extraction failed")
		}
	}
}
