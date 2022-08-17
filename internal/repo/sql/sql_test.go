package sql

import (
	"context"
	"testing"

	// 3rd party
	"github.com/stretchr/testify/require"

	// internal
	"github.com/TonyPath/user-mng-grpc-service/dockertest"
)

func TestNewDB(t *testing.T) {
	ctx := context.TODO()

	envArgs := []string{
		"POSTGRES_USER=db_user",
		"POSTGRES_PASSWORD=db_pwd",
		"POSTGRES_DB=db_test",
	}

	teardown, pgHost, err := dockertest.SetupPostgres(envArgs)
	if err != nil {
		t.Fatal(err)
	}
	defer teardown()

	cfg := Config{
		Host:     pgHost,
		DBName:   "db_test",
		User:     "db_user",
		Password: "db_pwd",
	}

	db, err := NewDB(cfg)
	require.NoError(t, err)
	require.NoError(t, StatusCheck(ctx, db))
	require.NoError(t, db.Close())
}
