package serializeBench

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
)

func BenchmarkJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bt, err := json.Marshal(exampleObject)
		if err != nil {
			panic(err)
		}
		resultBytes = bt
	}
}

func BenchmarkGOBEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var bytesGOB bytes.Buffer
		enc := gob.NewEncoder(&bytesGOB)

		err := enc.Encode(exampleObject)
		if err != nil {
			panic(err)
		}
		resultBytes = bytesGOB.Bytes()
	}
}

func BenchmarkProtoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bt, err := proto.Marshal(&exampleObject)
		if err != nil {
			panic(err)
		}
		resultBytes = bt
	}
}

func BenchmarkProtoGogoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bt, err := gogoproto.Marshal(&exampleObjectGogo)
		if err != nil {
			panic(err)
		}
		resultBytes = bt
	}
}
