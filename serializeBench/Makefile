.PHONY: proto all

proto:
	protoc pb/test.proto --go_out=.
	protoc pbgogo/test.proto --gofast_out=.

test:
	go test -v -bench=. -benchmem .