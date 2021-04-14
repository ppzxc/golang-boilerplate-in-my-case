package mariadb

import (
	"context"
	"database/sql"
	"github.com/ppzxc/golang-boilerplate-in-my-case/domain"
	"go.uber.org/zap"
)

type mariadbUserRepository struct {
	Conn *sql.DB
}

func NewMariadbUserRepository(conn *sql.DB) domain.UserRepository {
	return &mariadbUserRepository{
		Conn: conn,
	}
}

func (m *mariadbUserRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	query := `SELECT id, name, username, password, email, description, register_date, last_login_date `
	query += `FROM users `
	query += `WHERE id = ?`

	row := m.Conn.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		zap.L().Error("user.GetByID query Error",
			zap.String("query", query),
			zap.Error(row.Err()))
		return domain.User{}, row.Err()
	}

	user := domain.User{}
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Description,
		&user.RegisterDate,
		&user.LastLoginDate,
	); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Info("db.GetByID, rows scan not found",
				zap.Int64("ID", id),
				zap.Error(err))
			return domain.User{}, err
		} else {
			zap.L().Error("rows scan error",
				zap.Error(err))
			return domain.User{}, err
		}
	} else {
		return user, nil
	}
}

func (m *mariadbUserRepository) GetByEmail(ctx context.Context, email string) (res domain.User, err error) {
	query := `SELECT id, name, username, password, email, description, register_date, last_login_date `
	query += `FROM users `
	query += `WHERE email = ?`

	row := m.Conn.QueryRowContext(ctx, query, email)
	if row.Err() != nil {
		zap.L().Error("user.GetByID query Error",
			zap.String("query", query),
			zap.Error(row.Err()))
		return domain.User{}, row.Err()
	}

	user := domain.User{}
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Description,
		&user.RegisterDate,
		&user.LastLoginDate,
	); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Info("db.GetByEmail, rows scan not found",
				zap.String("email", email),
				zap.Error(err))
			return domain.User{}, err
		} else {
			zap.L().Error("rows scan error",
				zap.Error(err))
			return domain.User{}, err
		}
	} else {
		return user, nil
	}
}

func (m *mariadbUserRepository) Update(ctx context.Context, user *domain.User) error {
	return nil
}

func (m *mariadbUserRepository) Store(ctx context.Context, user *domain.User) error {
	return nil
}

func (m *mariadbUserRepository) Delete(ctx context.Context, id int64) error {
	return nil
}
