# Protobuf `proto.Marshall` vs `protojson.Marshall`

## Test Datatype

For this test we used `TickData` from `go-archiver`.

`qubic.proto:`
```proto
syntax = "proto3";  
  
option go_package = "/pb";  
  
message TickData {  
  uint32 computor_index = 1;  
  uint32 epoch = 2;  
  uint32 tick_number = 3;  
  uint64 timestamp = 4;  
  bytes var_struct = 5;  
  bytes time_lock = 6;  
  repeated string transaction_ids = 7;  
  repeated int64 contract_fees = 8;  
  string signature_hex = 9;  
}
```

## Benchmark code
```go
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
```

## Test Results

### Size
`protojson.Marshall()` was found to produce, on average, output `1.7x` to sometimes `2x` larger than `proto.Marshall()`.
In this test case the output is `2.07` times larger.

### Time
`protojson.Marshall()` seems to take on average about `10x` more time to serialize than `proto.Marshall().`

### Benchmark
`go test -bench=. -benchtime=10000x`

```
goos: linux
goarch: amd64
pkg: github.com/linckode/go-protobuf-marshall-test
cpu: Intel(R) Core(TM) i5-3570K CPU @ 3.40GHz
BenchmarkMarshallSpeed-4       	   10000	      3273 ns/op
--- BENCH: BenchmarkMarshallSpeed-4
    qubic_test.go:27: serialized message length: 559
    qubic_test.go:27: serialized message length: 559
BenchmarkJsonMarshallSpeed-4   	   10000	     32386 ns/op
--- BENCH: BenchmarkJsonMarshallSpeed-4
    qubic_test.go:40: json serialized message length: 1158
    qubic_test.go:40: json serialized message length: 1158
PASS
ok  	github.com/linckode/go-protobuf-marshall-test	0.674s

```
## Conclusion
In conclusion the Json variant for serializing is ~10 times slower while producing an output 1.6 to 2 times larger.
Json, as a format is based on text, thus it takes more space to represent data when compared to raw binary.

If performance and output size are a concern, `proto.Marshal()` is the recommended option.
`protojson.Marshal()` does have it's uses too, for example when replying to a http request.