# Serialization benchmark

Testing packages:
* Standard JSON package (encoding/json)
* Standard gob package (encoding/gob)
* Golang protobuf [https://github.com/golang/protobuf](github.com/golang/protobuf/proto)
* Gogo protobuf [https://github.com/gogo/protobuf](github.com/gogo/protobuf/proto)

```bash
$ make test
go test -v -bench=. -benchmem .
=== RUN   TestSerializeDeserializeJSON
--- PASS: TestSerializeDeserializeJSON (0.00s)
    serialize_test.go:23: JSON serialized string length = 256
=== RUN   TestSerializeDeserializeGOB
--- PASS: TestSerializeDeserializeGOB (0.00s)
    serialize_test.go:43: GOB serialized string length = 487
=== RUN   TestSerializeDeserializeProto
--- PASS: TestSerializeDeserializeProto (0.00s)
    serialize_test.go:63: Proto serialized string length = 108
=== RUN   TestSerializeDeserializeProtoGogo
--- PASS: TestSerializeDeserializeProtoGogo (0.00s)
    serialize_test.go:80: Proto (gogo) serialized string length = 108
goos: darwin
goarch: amd64
pkg: github.com/reinventer/go-samples/serializeBench
BenchmarkJSONMarshal-8                    470976              2353 ns/op             824 B/op         11 allocs/op
BenchmarkGOBEncode-8                       88328             13553 ns/op            3608 B/op         56 allocs/op
BenchmarkProtoMarshal-8                   546747              2123 ns/op             704 B/op         23 allocs/op
BenchmarkProtoGogoMarshal-8              2829918               419 ns/op             112 B/op          1 allocs/op
BenchmarkJSONUnmarshal-8                  166059              7391 ns/op            1128 B/op         37 allocs/op
BenchmarkGOBDecode-8                       31486             38732 ns/op           11701 B/op        314 allocs/op
BenchmarkProtoUnmarshal-8                 518366              2244 ns/op             808 B/op         27 allocs/op
BenchmarkProtoGogoUnmarshal-8            1000000              1014 ns/op             760 B/op         21 allocs/op
PASS
ok      github.com/reinventer/go-samples/serializeBench 10.414s
```