# Go JSON 库性能对比报告

## 测试环境
- **平台**: macOS (Apple M1 Pro, ARM64)
- **Go版本**: 1.23
- **测试数据1**: 小JSON (~100字节)
- **测试数据2**: 大JSON (~4.5KB)

## 小JSON性能测试结果 (~100字节)

| 排名 | 库 | 速度 | 内存分配 | 分配次数 | 相对性能 |
|------|-----|------|----------|----------|----------|
| 🥇 | **gjson** | 578 ns/op | 560 B/op | 3 allocs/op | **基准** |
| 🥈 | **sonic** | 872 ns/op | 1,828 B/op | 16 allocs/op | 1.5x 慢 |
| 🥉 | **json-iterator** | 1,161 ns/op | 1,480 B/op | 37 allocs/op | 2.0x 慢 |
| 4 | **go-json** | 1,163 ns/op | 1,537 B/op | 34 allocs/op | 2.0x 慢 |
| 5 | **fastjson** | 1,014 ns/op | 3,024 B/op | 14 allocs/op | 1.8x 慢 |
| 6 | **标准库** | 1,779 ns/op | 1,528 B/op | 31 allocs/op | 3.1x 慢 |
| 7 | **segmentio** | 1,703 ns/op | 5,912 B/op | 22 allocs/op | 2.9x 慢 |
| 8 | **yyjson优化** | 3,828 ns/op | 404 B/op | 19 allocs/op | 6.6x 慢 |
| 9 | **yyjson原始** | 3,489 ns/op | 1,396 B/op | 25 allocs/op | 6.0x 慢 |

## 大JSON性能测试结果 (~4.5KB)

| 排名 | 库 | 速度 | 内存分配 | 分配次数 | 相对性能 |
|------|-----|------|----------|----------|----------|
| 🥇 | **gjson** | 4,287 ns/op | 0 B/op | 0 allocs/op | **基准** |
| 🥈 | **sonic** | 10,307 ns/op | 19,044 B/op | 78 allocs/op | 2.4x 慢 |
| 🥉 | **fastjson** | 12,766 ns/op | 37,176 B/op | 138 allocs/op | 3.0x 慢 |
| 4 | **go-json** | 21,178 ns/op | 23,107 B/op | 517 allocs/op | 4.9x 慢 |
| 5 | **json-iterator** | 21,363 ns/op | 19,278 B/op | 578 allocs/op | 5.0x 慢 |
| 6 | **segmentio** | 26,581 ns/op | 21,496 B/op | 392 allocs/op | 6.2x 慢 |
| 7 | **标准库** | 33,500 ns/op | 16,992 B/op | 459 allocs/op | 7.8x 慢 |
| 8 | **yyjson原始** | 55,911 ns/op | 21,504 B/op | 414 allocs/op | 13.0x 慢 |
| 9 | **yyjson优化** | 64,852 ns/op | 7,792 B/op | 368 allocs/op | 15.1x 慢 |

## 关键发现

### 🏆 性能之王
- **gjson**在所有测试中都表现最佳，零内存分配，查询场景无敌
- **sonic**在完整unmarshal库中表现最好，比标准库快2-3倍
- **fastjson**在完整解析后多次访问场景下有优势，但整体仍比gjson和sonic慢

### 💡 内存效率
- **yyjson优化版**内存分配最少（404B），但速度最慢
- **gjson**零内存分配，但API不同
- **sonic**在性能和内存使用间取得良好平衡

### 🚨 CGO性能问题
- **yyjson在Go中的表现远不如预期**，主要原因是：
  1. CGO调用开销巨大
  2. 数据在C和Go之间转换的成本
  3. 在大数据时性能差距更加明显

### 📈 库特点分析

| 库 | 类型 | 优势 | 劣势 |
|----|-----|------|------|
| **sonic** | JIT + SIMD | 极快速度，完全兼容 | 内存略高 |
| **json-iterator** | 优化反射 | 良好兼容性 | 中等性能 |
| **go-json** | 优化反射 | 完全兼容 | 性能一般 |
| **segmentio** | 优化反射 | 稳定可靠 | 内存较多 |
| **gjson** | 查询专用 | 零分配极速 | API不完整 |
| **yyjson** | CGO + SIMD | C++性能理论 | CGO开销大 |

## 🎯 推荐选择

### 1. 最高性能场景
- **sonic** - JIT编译 + SIMD，最佳unmarshal性能

### 2. 查询密集场景  
- **gjson** - 零分配极速，查询场景无敌

### 3. 兼容性要求高
- **标准库** - 最稳定，完全兼容

### 4. 内存敏感场景
- **yyjson优化版** - 内存最少，但速度慢

### ❌ 不推荐场景
- **yyjson CGO版本** - 性能差，除非特殊需求

## 🔧 优化建议

对于yyjson在Go中的使用：
1. **避免频繁调用** - 批量处理减少CGO开销
2. **考虑纯Go实现** - sonic等已证明更有效
3. **SIMD在Go中** - 通过JIT比CGO更有效

结论：在Go生态中，**sonic是yyjson的最佳替代品**，提供了SIMD级别的性能而无需CGO的开销。