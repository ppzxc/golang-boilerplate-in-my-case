package domain

import (
	"context"
	"time"
)

// User ...
type User struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Username      string    `json:"username" validate:"required"`
	Password      string    `json:"password" validate:"required"`
	Email         string    `json:"email" validate:"required,email"`
	Description   string    `json:"description"`
	RegisterDate  time.Time `json:"register_date"`
	LastLoginDate time.Time `json:"last_login_date"`
}

// UserUsecase represent the user's repository contract
type UserUsecase interface {
	GetByID(ctx context.Context, id int64) (User, error)
	GetByEmail(ctx context.Context, Email string) (User, error)
	Update(ctx context.Context, user *User) error
	Store(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	GetByID(ctx context.Context, id int64) (User, error)
	GetByEmail(ctx context.Context, Email string) (User, error)
	Update(ctx context.Context, user *User) error
	Store(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}
