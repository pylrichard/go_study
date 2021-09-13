package add

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// BenchmarkAddStringWithOperator 字符串是不可变的，每次会产生临时新字符串，给gc造成额外负担
func BenchmarkAddStringWithOperator(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = hello + "=" + world
	}
}

// BenchmarkAddStringWithSprintf 格式化字符串，拼接数字
func BenchmarkAddStringWithSprintf(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s=%s", hello, world)
	}
}

// BenchmarkAddStringWithJoin join()会先根据字符串数组的内容计算出一个拼接之后的长度
// 然后申请对应大小的内存，一个一个字符串填入，在已有一个数组的情况下，这种效率会很高
func BenchmarkAddStringWithJoin(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{hello, world}, "=")
	}
}

// BenchmarkAddStringWithBuffer 可以当可变字符使用，对内存的增长也有优化
// 如果能预估字符串长度，还可以用buffer.Grow()接口设置capacity
func BenchmarkAddStringWithBuffer(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(hello)
		buffer.WriteString("=")
		buffer.WriteString(world)
		_ = buffer.String()
	}
}