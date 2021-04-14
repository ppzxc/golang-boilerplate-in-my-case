package domain

type Auth struct {
	Token string `json:"token" validate:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

//// AuthUsecase represent the user's repository contract
//type AuthUsecase interface {
//	Login(ctx context.Context, Email string) (User, error)
//}

//// AuthRepository represent the user's repository contract
//type AuthRepository interface {
//	GetByEmail(ctx context.Context, Email string) (User, error)
//}
