package serializeBench

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"

	"github.com/reinventer/go-samples/serializeBench/pb"
	"github.com/reinventer/go-samples/serializeBench/pbgogo"
)

func TestSerializeDeserializeJSON(t *testing.T) {
	bytesJSON, err := json.Marshal(exampleObject)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("JSON serialized string length = %d", len(bytesJSON))

	var res pb.TestObject
	err = json.Unmarshal(bytesJSON, &res)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, exampleObject, res)
}

func TestSerializeDeserializeGOB(t *testing.T) {
	var bytesGOB bytes.Buffer
	enc := gob.NewEncoder(&bytesGOB)

	err := enc.Encode(exampleObject)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GOB serialized string length = %d", bytesGOB.Len())

	var res pb.TestObject
	dec := gob.NewDecoder(&bytesGOB)
	err = dec.Decode(&res)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, exampleObject, res)
}

func TestSerializeDeserializeProto(t *testing.T) {
	sourceObject := proto.Clone(&exampleObject)

	bytesProto, err := proto.Marshal(sourceObject)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Proto serialized string length = %d", len(bytesProto))

	var res pb.TestObject
	err = proto.Unmarshal(bytesProto, &res)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, exampleObject, res)
}

func TestSerializeDeserializeProtoGogo(t *testing.T) {
	bytesProto, err := gogoproto.Marshal(&exampleObjectGogo)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Proto (gogo) serialized string length = %d", len(bytesProto))

	var res pbgogo.TestObject
	err = gogoproto.Unmarshal(bytesProto, &res)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, exampleObjectGogo, res)
}
