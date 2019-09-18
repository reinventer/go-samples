package serializeBench

import (
	"github.com/reinventer/go-samples/serializeBench/pb"
	"github.com/reinventer/go-samples/serializeBench/pbgogo"
)

var exampleObject = pb.TestObject{
	Query:         "query1",
	PageNumber:    1,
	ResultPerPage: 2,
	Flag:          true,
	Score:         3.4,
	List: []*pb.TestNested{
		{Title: "nested1"},
		{Title: "nested2"},
		{Title: "nested3"},
	},
	Map: map[int64]*pb.TestNested{
		4: {Title: "nested4"},
		5: {Title: "nested5"},
		6: {Title: "nested6"},
	},
	Nested: &pb.TestNested{
		Title: "nested7",
	},
}

var exampleObjectGogo = pbgogo.TestObject{
	Query:         "query1",
	PageNumber:    1,
	ResultPerPage: 2,
	Flag:          true,
	Score:         3.4,
	List: []*pbgogo.TestNested{
		{Title: "nested1"},
		{Title: "nested2"},
		{Title: "nested3"},
	},
	Map: map[int64]*pbgogo.TestNested{
		4: {Title: "nested4"},
		5: {Title: "nested5"},
		6: {Title: "nested6"},
	},
	Nested: &pbgogo.TestNested{
		Title: "nested7",
	},
}

var resultBytes []byte
var resultObject *pb.TestObject
var resultObjectGogo *pbgogo.TestObject
