package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	// 3rd party
	"github.com/google/uuid"
	"go.uber.org/zap"

	// internal
	"github.com/TonyPath/user-mng-grpc-service/internal/models"
	pg "github.com/TonyPath/user-mng-grpc-service/internal/repo/sql"
)

const usersTable = "users"

type Repository struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewRepository(db *sql.DB, log *zap.SugaredLogger) *Repository {
	return &Repository{
		db:     db,
		logger: log,
	}
}

func (r *Repository) InsertUser(ctx context.Context, user models.User) (uuid.UUID, error) {
	query, args, err := pg.QueryBuilder().
		Insert(usersTable).
		Columns("id", "email", "first_name", "last_name", "nickname", "password", "country").
		Values(user.ID, user.Email, user.FirstName, user.LastName, user.Nickname, user.Password, user.Country).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return uuid.Nil, nil
	}

	row := r.db.QueryRowContext(ctx, query, args...)

	var userID uuid.UUID
	err = row.Scan(&userID)
	if err != nil {
		if pg.IsUniqueViolation(err) {
			return uuid.Nil, models.ErrEmailTaken
		}
		return uuid.Nil, err
	}

	return userID, nil
}

func (r *Repository) UpdateUser(ctx context.Context, userID uuid.UUID, user models.User) error {
	query, args, err := pg.QueryBuilder().
		Update(usersTable).
		Set("email", user.Email).
		Set("first_name", user.FirstName).
		Set("last_name", user.LastName).
		Set("nickname", user.Nickname).
		Set("password", user.Password).
		Set("country", user.Country).
		Set("updated_at", user.UpdateAt).
		Where("id = ?", userID).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		if pg.IsUniqueViolation(err) {
			return models.ErrEmailTaken
		}
		return err
	}

	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	query, args, err := pg.QueryBuilder().
		Delete(usersTable).
		Where("id = ?", userID).
		ToSql()

	if err != nil {
		return fmt.Errorf("could not build query sql query: %w", err)
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUsersByFilter(ctx context.Context, opts models.GetUsersOptions) ([]models.User, error) {
	qb := pg.QueryBuilder().
		Select("id", "email", "first_name", "last_name", "nickname", "password", "country", "created_at", "updated_at").
		From(usersTable).
		Suffix("OFFSET ? ROWS FETCH NEXT ? ROWS ONLY", (opts.PageNumber-1)*opts.PageSize, opts.PageSize)

	if opts.Filter.Country != "" {
		qb = qb.Where("country = ?", opts.Filter.Country)
	}

	if opts.Filter.Email != "" {
		qb = qb.Where("email = ?", opts.Filter.Email)
	}

	if opts.Filter.Nickname != "" {
		qb = qb.Where("nickname = ?", opts.Filter.Nickname)
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(
			&u.ID,
			&u.Email,
			&u.FirstName,
			&u.LastName,
			&u.Nickname,
			&u.Password,
			&u.Country,
			&u.CreatedAt,
			&u.UpdateAt,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	qb := pg.QueryBuilder().
		Select("id", "email", "first_name", "last_name", "nickname", "password", "country", "created_at", "updated_at").
		From(usersTable).
		Where("id = ?", userID)

	query, args, err := qb.ToSql()
	if err != nil {
		return models.User{}, err
	}

	var u models.User
	row := r.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&u.ID,
		&u.Email,
		&u.FirstName,
		&u.LastName,
		&u.Nickname,
		&u.Password,
		&u.Country,
		&u.CreatedAt,
		&u.UpdateAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, models.ErrUserNotFound
		}
		return models.User{}, err
	}

	return u, nil
}

func (r *Repository) ExistsByID(ctx context.Context, userID uuid.UUID) (bool, error) {
	query, args, err := pg.QueryBuilder().
		Select("1").
		Prefix("SELECT EXISTS(").
		From(usersTable).
		Where("id = ?", userID).
		Limit(1).
		Suffix(")").
		ToSql()

	if err != nil {
		return false, fmt.Errorf("could not build query sql query: %w", err)
	}

	var exists bool
	row := r.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("could not query select: %w", err)
	}

	return exists, nil
}
