package usecase

import (
	"context"
	"github.com/ppzxc/golang-boilerplate-in-my-case/domain"
	"time"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) GetByID(ctx context.Context, id int64) (res domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if res, err = u.userRepo.GetByID(ctx, id); err != nil {
		return
	}

	return
}

func (u *userUsecase) GetByEmail(ctx context.Context, email string) (res domain.User, err error) {
	return
}

func (u *userUsecase) Update(ctx context.Context, user *domain.User) error {
	return nil
}

func (u *userUsecase) Store(ctx context.Context, user *domain.User) error {
	return nil
}

func (u *userUsecase) Delete(ctx context.Context, id int64) error {
	return nil
}
