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

func (m *mariadbUserRepository) GetByID(ctx context.Context, id int64) (res domain.User, err error) {
	query := `SELECT id, name, username, password, email, description, register_date, last_login_date `
	query += `FROM users `
	query += `WHERE id = ?`

	rows, err := m.Conn.QueryContext(ctx, query, id)
	if err != nil {
		zap.L().Error("user.GetByID query Error",
			zap.String("query", query),
			zap.Error(err))
		return domain.User{}, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			zap.L().Error("user.GetByID rows close err",
				zap.String("query", query),
				zap.Error(err))
		}
	}()

	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.Description,
			&user.RegisterDate,
			&user.LastLoginDate,
		)

		if err != nil {
			zap.L().Error("rows scan error", zap.Error(err))
			return domain.User{}, err
		}

		res = user
	}

	return
}

func (m *mariadbUserRepository) GetByEmail(ctx context.Context, email string) (res domain.User, err error) {
	return
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
