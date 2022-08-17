package sql

import (
	"context"
	"database/sql"
	"net/url"
	"time"

	// 3rd party
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	DBName   string
	User     string
	Password string
}

func NewDB(cfg Config) (*sql.DB, error) {
	q := make(url.Values)
	q.Set("sslmode", "disable")
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host,
		Path:     cfg.DBName,
		RawQuery: q.Encode(),
	}

	db, err := sql.Open("postgres", u.String())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func StatusCheck(ctx context.Context, db *sql.DB) error {
	var (
		pingMaxAttempts = 20
		pingError       error
	)

	for attempts := 1; ; attempts++ {
		pingError = db.Ping()
		if pingError == nil {
			break
		}
		if attempts == pingMaxAttempts {
			return pingError
		}
		time.Sleep(time.Duration(attempts) * 10 * time.Millisecond)
	}

	const query = `SELECT true`
	var tmp bool
	return db.QueryRowContext(ctx, query).Scan(&tmp)
}

func QueryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
