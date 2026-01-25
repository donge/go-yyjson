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

// 生成约2000字节的复杂JSON
var largeJSONData = []byte(`{
	"user": {
		"id": 12345,
		"name": "John Doe",
		"email": "john.doe@example.com",
		"avatar": "https://example.com/avatar.jpg",
		"verified": true,
		"created_at": "2023-01-15T10:30:00Z",
		"updated_at": "2023-12-20T15:45:30Z",
		"settings": {
			"theme": "dark",
			"notifications": true,
			"privacy": {
				"show_email": false,
				"show_phone": true,
				"allow_search": true
			},
			"preferences": {
				"language": "en",
				"timezone": "UTC",
				"currency": "USD",
				"date_format": "YYYY-MM-DD"
			}
		},
		"roles": ["user", "moderator", "admin"],
		"permissions": {
			"read": true,
			"write": true,
			"delete": false,
			"admin": true
		}
	},
	"posts": [
		{
			"id": 1,
			"title": "First Post",
			"content": "This is the content of the first post with some detailed text to make it longer.",
			"tags": ["intro", "welcome", "first"],
			"metadata": {
				"views": 1250,
				"likes": 89,
				"comments": 23,
				"shares": 12,
				"published": true,
				"featured": false
			},
			"author": {
				"id": 12345,
				"name": "John Doe"
			},
			"created_at": "2023-06-01T12:00:00Z"
		},
		{
			"id": 2,
			"title": "Second Post",
			"content": "This is the second post with even more content to increase the JSON size significantly.",
			"tags": ["update", "announcement", "important"],
			"metadata": {
				"views": 3420,
				"likes": 256,
				"comments": 67,
				"shares": 45,
				"published": true,
				"featured": true
			},
			"author": {
				"id": 12345,
				"name": "John Doe"
			},
			"created_at": "2023-07-15T14:30:00Z"
		},
		{
			"id": 3,
			"title": "Third Post",
			"content": "Another post with substantial content to make the JSON data larger for performance testing purposes.",
			"tags": ["tutorial", "guide", "learning"],
			"metadata": {
				"views": 2180,
				"likes": 145,
				"comments": 34,
				"shares": 28,
				"published": true,
				"featured": false
			},
			"author": {
				"id": 12345,
				"name": "John Doe"
			},
			"created_at": "2023-08-20T16:45:00Z"
		}
	],
	"analytics": {
		"total_posts": 3,
		"total_views": 6850,
		"total_likes": 490,
		"total_comments": 124,
		"total_shares": 85,
		"engagement_rate": 0.072,
		"avg_views_per_post": 2283.33,
		"avg_likes_per_post": 163.33,
		"top_performing_post": {
			"id": 2,
			"title": "Second Post",
			"views": 3420,
			"engagement_score": 0.092
		},
		"daily_stats": {
			"monday": {"posts": 0, "views": 120, "likes": 8},
			"tuesday": {"posts": 1, "views": 450, "likes": 32},
			"wednesday": {"posts": 0, "views": 280, "likes": 19},
			"thursday": {"posts": 1, "views": 890, "likes": 67},
			"friday": {"posts": 1, "views": 2340, "likes": 178},
			"saturday": {"posts": 0, "views": 1560, "likes": 98},
			"sunday": {"posts": 0, "views": 1210, "likes": 88}
		}
	},
	"recommendations": [
		{
			"type": "user",
			"id": 67890,
			"name": "Jane Smith",
			"reason": "similar interests",
			"score": 0.85
		},
		{
			"type": "post",
			"id": 456,
			"title": "Advanced Techniques",
			"reason": "based on your reading history",
			"score": 0.92
		},
		{
			"type": "topic",
			"name": "Technology",
			"reason": "trending in your network",
			"score": 0.78
		}
	],
	"notifications": {
		"unread_count": 5,
		"total_count": 23,
		"settings": {
			"email_notifications": true,
			"push_notifications": true,
			"sms_notifications": false,
			"weekly_digest": true
		},
		"recent": [
			{
				"id": 1001,
				"type": "like",
				"message": "Jane liked your post",
				"timestamp": "2023-12-20T10:15:00Z",
				"read": false
			},
			{
				"id": 1002,
				"type": "comment",
				"message": "Bob commented on your post",
				"timestamp": "2023-12-20T09:30:00Z",
				"read": false
			}
		]
	}
}`)

func BenchmarkLargeUnmarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data map[string]interface{}
		if err := Unmarshal(largeJSONData, &data); err != nil {
			b.Fatalf("Unmarshal error: %v", err)
		}
	}
}

func BenchmarkLargeStdUnmarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data map[string]interface{}
		if err := json.Unmarshal(largeJSONData, &data); err != nil {
			b.Fatalf("json.Unmarshal error: %v", err)
		}
	}
}

func BenchmarkLargeUnmarshalOptimized(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data map[string]interface{}
		if err := UnmarshalOptimized(largeJSONData, &data); err != nil {
			b.Fatalf("UnmarshalOptimized error: %v", err)
		}
	}
}

func BenchmarkLargeSonicUnmarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data map[string]interface{}
		if err := sonic.Unmarshal(largeJSONData, &data); err != nil {
			b.Fatalf("sonic.Unmarshal error: %v", err)
		}
	}
}

func BenchmarkLargeGjsonParse(b *testing.B) {
	jsonStr := string(largeJSONData)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 使用 gjson 解析大JSON并获取多个值来模拟完整使用
		result := gjson.Parse(jsonStr)

		// 获取主要的嵌套值来模拟实际使用场景
		userName := result.Get("user.name").String()
		totalViews := result.Get("analytics.total_views").Int()
		postCount := result.Get("analytics.total_posts").Int()
		topPostViews := result.Get("analytics.top_performing_post.views").Int()
		unreadCount := result.Get("notifications.unread_count").Int()

		// 验证解析是否成功
		if !result.Exists() || userName == "" || totalViews == 0 || postCount == 0 || topPostViews == 0 || unreadCount == 0 {
			b.Fatalf("gjson.Parse or value extraction failed")
		}
	}
}

func BenchmarkLargeJsoniterUnmarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data map[string]interface{}
		if err := jsoniter.Unmarshal(largeJSONData, &data); err != nil {
			b.Fatalf("jsoniter.Unmarshal error: %v", err)
		}
	}
}

func BenchmarkLargeGojsonUnmarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data map[string]interface{}
		if err := gojson.Unmarshal(largeJSONData, &data); err != nil {
			b.Fatalf("gojson.Unmarshal error: %v", err)
		}
	}
}

func BenchmarkLargeSegmentjsonUnmarshal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var data map[string]interface{}
		if err := segmentjson.Unmarshal(largeJSONData, &data); err != nil {
			b.Fatalf("segmentjson.Unmarshal error: %v", err)
		}
	}
}
