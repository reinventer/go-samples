# Benchmark for postgres drivers

```
$ go test -bench . -benchmem github.com/reinventer/go-samples/pgbench
testing: warning: no tests to run
BenchmarkPq-4        	   50000	     29861 ns/op	     417 B/op	      15 allocs/op
BenchmarkPg-4        	   50000	     30454 ns/op	     113 B/op	       5 allocs/op
BenchmarkPgxSql-4    	   30000	     54708 ns/op	     587 B/op	      14 allocs/op
BenchmarkPgxPure-4   	   30000	     54439 ns/op	     364 B/op	       4 allocs/op
PASS
ok  	github.com/reinventer/go-samples/pgbench	8.242s
```

BenchmarkPq - [github.com/lib/pq](https://github.com/lib/pq/) with `database/sql` interface
BenchmarkPg - [github.com/go-pg/pg](https://github.com/go-pg/pg) ORM for Golang
BenchmarkPgxSql - [github.com/jackc/pgx](https://github.com/jackc/pgx) with `database/sql` interface
BenchmarkPgxPure - pure [github.com/jackc/pgx](https://github.com/jackc/pgx)
