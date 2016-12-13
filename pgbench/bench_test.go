package pgbench

import (
	"database/sql"
	"flag"
	"os/user"
	"strconv"
	"strings"
	"testing"

	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"gopkg.in/pg.v5"
)

const query = "SELECT 1"

var (
	num                                int
	err                                error
	pgbase, pguser, pgpassword, pghost string
	pgport                             int
)

func init() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	flag.StringVar(&pgbase, "pgbase", "postgres", "postgres database")
	flag.StringVar(&pguser, "pguser", u.Username, "postgres user")
	flag.StringVar(&pgpassword, "pgpassword", "", "postgres password")
	flag.StringVar(&pghost, "pghost", "localhost", "postgres host")
	flag.IntVar(&pgport, "pgport", 5432, "postgres port")
	flag.Parse()
}

func BenchmarkPq(b *testing.B) {
	db, err := sql.Open("postgres", sqlURI())
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(5)

	stmt, err := db.Prepare(query)
	if err != nil {
		b.Fatal(err)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err = stmt.QueryRow().Scan(&num); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPg(b *testing.B) {
	db := pg.Connect(&pg.Options{
		User:     pguser,
		Password: pgpassword,
		Database: pgbase,
		Addr:     strings.Join([]string{pghost, ":", strconv.Itoa(pgport)}, ""),
	})
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if _, err = db.Query(pg.Scan(&num), query); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPgxSql(b *testing.B) {
	db, err := sql.Open("pgx", sqlURI())
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(5)

	stmt, err := db.Prepare(query)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err = stmt.QueryRow().Scan(&num); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkPgxPure(b *testing.B) {
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Database: pgbase,
			Host:     pghost,
			Port:     uint16(pgport),
			User:     pguser,
			Password: pgpassword,
		},
		MaxConnections: 15,
	})
	if err != nil {
		b.Fatal(err)
	}
	defer pool.Close()

	_, err = pool.Prepare("testquery", query)
	if err != nil {
		b.Fatal(err)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err = pool.QueryRow("testquery").Scan(&num); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func sqlURI() string {
	return strings.Join([]string{"postgres://", pguser, ":", pgpassword, "@", pghost, ":", strconv.Itoa(pgport), "/", pgbase, "?sslmode=disable"}, "")
}
