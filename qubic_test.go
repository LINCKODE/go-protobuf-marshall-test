package main

import (
	"github.com/linckode/go-protobuf-marshall-test/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"testing"
)

var (
	message = &pb.TickData{
		Timestamp:      uint64(1),
		TickNumber:     uint32(2),
		ComputorIndex:  uint32(3),
		Epoch:          uint32(4),
		ContractFees:   make([]int64, 100),
		SignatureHex:   "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		TimeLock:       make([]byte, 100),
		TransactionIds: make([]string, 100),
		VarStruct:      make([]byte, 100),
	}
)

func BenchmarkMarshallSpeed(b *testing.B) {

	data, _ := proto.Marshal(message)
	b.Logf("serialized message length: %d", len(data))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		_, _ = proto.Marshal(message)
		b.StopTimer()
	}
}

func BenchmarkJsonMarshallSpeed(b *testing.B) {

	data, _ := protojson.Marshal(message)
	b.Logf("json serialized message length: %d", len(data))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		_, _ = protojson.Marshal(message)
		b.StopTimer()
	}
}
