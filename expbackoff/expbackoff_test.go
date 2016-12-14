package expbackoff

import (
	"testing"
	"time"
)

// $ go test -bench . github.com/reinventer/go-samples/expbackoff
// BenchmarkNoJam-4     3000000           522 ns/op
// PASS
// ok      github.com/reinventer/go-samples/expbackoff 2.109s
func BenchmarkNoJam(b *testing.B) {
	bo := New(time.Microsecond, time.Minute)

	c := bo.Fetch()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			<-c
		}
	})

	// bo.Stop()
}
