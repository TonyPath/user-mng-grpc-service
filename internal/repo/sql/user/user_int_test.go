package user

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	// 3rd party
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	// internal
	"github.com/TonyPath/user-mng-grpc-service/dockertest"
	"github.com/TonyPath/user-mng-grpc-service/internal/models"
	"github.com/TonyPath/user-mng-grpc-service/internal/repo/sql"
	"github.com/TonyPath/user-mng-grpc-service/internal/repo/sql/sqltest"
)

var (
	testDB *sqltest.DB
)

func TestMain(m *testing.M) {
	exitCode := run(m)
	os.Exit(exitCode)
}

func run(m *testing.M) int {

	envArgs := []string{
		"POSTGRES_USER=db_user",
		"POSTGRES_PASSWORD=db_pwd",
		"POSTGRES_DB=db_test",
	}

	teardown, pgHost, err := dockertest.SetupPostgres(envArgs)
	if err != nil {
		panic(err)
	}
	defer teardown()

	cfg := sql.Config{
		Host:     pgHost,
		DBName:   "db_test",
		User:     "db_user",
		Password: "db_pwd",
	}
	db, err := sql.NewDB(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = sql.StatusCheck(context.TODO(), db)
	if err != nil {
		panic(err)
	}

	testDB = &sqltest.DB{
		Db: db,
	}

	err = testDB.RunMigrations("file://./../../../../migrations/sql")
	if err != nil {
		panic(err)
	}

	return m.Run()
}

func TestRepository_InsertUser(t *testing.T) {

	repo := NewRepository(testDB.Db, zap.NewNop().Sugar())

	for i := 1; i <= 15; i++ {
		userID := uuid.New()
		user := models.User{
			ID:        userID,
			Email:     fmt.Sprintf("antonis+%d@mail.com", i),
			FirstName: fmt.Sprintf("tony+%d", i),
			LastName:  fmt.Sprintf("papath+%d", i),
			Nickname:  fmt.Sprintf("TonyPath+%d", i),
			Country:   "GR",
			Password:  []byte(`secret`),
			CreatedAt: time.Now(),
			UpdateAt:  nil,
		}

		gotUserID, err := repo.InsertUser(context.TODO(), user)
		require.NoError(t, err)
		require.Equal(t, userID, gotUserID)

		testDB.RequireTotalRows(t, "users", i)
	}
}

func TestRepository_GetUsersByFilter(t *testing.T) {

	repo := NewRepository(testDB.Db, zap.NewNop().Sugar())

	t.Log("1st page")
	{
		gotUsers, err := repo.GetUsersByFilter(context.TODO(), models.GetUsersOptions{
			PageNumber: 1,
			PageSize:   10,
		})
		require.NoError(t, err)
		require.Len(t, gotUsers, 10)
	}

	t.Log("2nd page")
	{
		gotUsers, err := repo.GetUsersByFilter(context.TODO(), models.GetUsersOptions{
			PageNumber: 2,
			PageSize:   10,
		})
		require.NoError(t, err)
		require.Len(t, gotUsers, 5)
	}
}

func TestRepository_DeleteUser(t *testing.T) {

	repo := NewRepository(testDB.Db, zap.NewNop().Sugar())

	users, err := repo.GetUsersByFilter(context.TODO(), models.GetUsersOptions{
		PageNumber: 1,
		PageSize:   1,
	})
	require.NoError(t, err)

	err = repo.DeleteUser(context.TODO(), users[0].ID)
	require.NoError(t, err)
	testDB.RequireTotalRows(t, "users", 14)
}
