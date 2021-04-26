package convert 

/*
	go test -bench="." -benchmem
	-benchmem可以提供每次操作分配内存的次数，以及每次操作分配的字节数

	强转换的性能明显优于标准转换
 */
import (
	"bytes"
	"testing"
)

/*
	测试强转换功能
 */
func TestBytes2Str(t *testing.T) {
	x := []byte("Hello Gopher!")
	y := bytes2Str(x)
	z := string(x)

	if y != z {
		t.Fail()
	}
}

func TestStr2Bytes(t *testing.T) {
	x := "Hello Gopher!"
	y := str2Bytes(x)
	z := []byte(x)

	if !bytes.Equal(y, z) {
		t.Fail()
	}
}

//测试标准转换string()性能
func BenchmarkNormalBytes2Str(b *testing.B) {
	x := []byte("Hello Gopher!Hello Gopher!Hello Gopher!")
	for i := 0; i < b.N; i++ {
		_ = string(x)
	}
}

//测试强转换[]byte到string性能
func BenchmarkBytes2Str(b *testing.B) {
	x := []byte("Hello Gopher!Hello Gopher!Hello Gopher!")
	for i := 0; i < b.N; i++ {
		_ = bytes2Str(x)
	}
}

//测试标准转换[]byte性能
func BenchmarkNormalStr2Bytes(b *testing.B) {
	x := "Hello Gopher! Hello Gopher! Hello Gopher!"
	for i := 0; i < b.N; i++ {
		_ = []byte(x)
	}
}

//测试强转换string到[]byte性能
func BenchmarkStr2Bytes(b *testing.B) {
	x := "Hello Gopher! Hello Gopher! Hello Gopher!"
	for i := 0; i < b.N; i++ {
		_ = str2Bytes(x)
	}
}