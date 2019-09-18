package serializeBench

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"

	"github.com/reinventer/go-samples/serializeBench/pb"
	"github.com/reinventer/go-samples/serializeBench/pbgogo"
)

func BenchmarkJSONUnmarshal(b *testing.B) {
	bt, err := json.Marshal(exampleObject)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var res pb.TestObject
		err := json.Unmarshal(bt, &res)
		if err != nil {
			panic(err)
		}
		resultObject = &res
	}
}

func BenchmarkGOBDecode(b *testing.B) {
	var bytesGOB bytes.Buffer
	enc := gob.NewEncoder(&bytesGOB)

	err := enc.Encode(exampleObject)
	if err != nil {
		panic(err)
	}

	bt := bytesGOB.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var res pb.TestObject
		dec := gob.NewDecoder(bytes.NewReader(bt))
		err := dec.Decode(&res)
		if err != nil {
			panic(err)
		}
		resultObject = &res
	}
}

func BenchmarkProtoUnmarshal(b *testing.B) {
	bt, err := proto.Marshal(&exampleObject)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var res pb.TestObject
		err := proto.Unmarshal(bt, &res)
		if err != nil {
			panic(err)
		}
		resultObject = &res
	}
}

func BenchmarkProtoGogoUnmarshal(b *testing.B) {
	bt, err := gogoproto.Marshal(&exampleObjectGogo)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var res pbgogo.TestObject
		err := gogoproto.Unmarshal(bt, &res)
		if err != nil {
			panic(err)
		}
		resultObjectGogo = &res
	}
}
