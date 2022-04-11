package language

import (
	"bytes"
	"encoding/json"
	"sync"
	"testing"
)

//go test -bench . -benchmem
/*
反序列化场景,使用pool降低GC压力
*/
var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}

//go test -bench="Buff*" . -benchmem
/*

 */
var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}
