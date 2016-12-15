# Exponential backoff implementation

About Full Jitter algorithm read [article](https://www.awsarchitectureblog.com/2015/03/backoff.html) at awsarchitectureblog.

Usage:

```sh
go get github.com/reinventer/go-samples/expbackoff
```

```go
package main

import "github.com/reinventer/go-samples/expbackoff"

func main() {
	b := expbackoff.New(
		2*time.Millisecond, // slot time
		10*time.Second, // maximum timeout
	)
	defer b.Stop()

	c := b.Fetch()
	for i := 0; i < 30; i++ {
		t := time.Now()
		<-c // pause in case of problem on backend

		// do some work

		if overloaded {
			b.Jam() // increase pause
		} else {
			b.Reset() // reset pause
		}
	}

}
```