# Benchmark for postgres drivers

```
$ go test -bench . -benchmem github.com/reinventer/go-samples/pgbench
testing: warning: no tests to run
BenchmarkPq-4        	   50000	     27539 ns/op	     401 B/op	      11 allocs/op
BenchmarkPg-4        	   50000	     28886 ns/op	     113 B/op	       5 allocs/op
BenchmarkPgxSql-4    	   50000	     27672 ns/op	     450 B/op	      11 allocs/op
BenchmarkPgxPure-4   	   50000	     25205 ns/op	     194 B/op	       0 allocs/op
PASS
ok  	github.com/reinventer/go-samples/pgbench	6.867s
```

BenchmarkPq - [github.com/lib/pq](https://github.com/lib/pq/) with `database/sql` interface

BenchmarkPg - [github.com/go-pg/pg](https://github.com/go-pg/pg) ORM for Golang

BenchmarkPgxSql - [github.com/jackc/pgx](https://github.com/jackc/pgx) with `database/sql` interface

BenchmarkPgxPure - pure [github.com/jackc/pgx](https://github.com/jackc/pgx)
