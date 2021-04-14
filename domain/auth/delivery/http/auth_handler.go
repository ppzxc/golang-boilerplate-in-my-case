package http

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/ppzxc/golang-boilerplate-in-my-case/domain"
	"github.com/ppzxc/golang-boilerplate-in-my-case/middleware"
	errUtil "github.com/ppzxc/golang-boilerplate-in-my-case/util/err"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	validator "gopkg.in/go-playground/validator.v9"
	"time"
)

type AuthHandler struct {
	//AuthUsecase domain.AuthUsecase
	UserUsecase domain.UserUsecase
	secret      string
}

//func NewAuthHandler(r fiber.Router, au domain.AuthUsecase, uu domain.UserUsecase) {

func NewAuthHandler(r fiber.Router, uu domain.UserUsecase, secret string) {
	handler := &AuthHandler{
		//AuthUsecase: au,
		UserUsecase: uu,
		secret:      secret,
	}

	user := r.Group("/auth")
	user.Post("/login", handler.Login)
	user.Post("/verify", middleware.Protected(secret), handler.Verify)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func isRequestValid(m *domain.Login) error {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return err
	}
	return nil
}

func (u *AuthHandler) Verify(c *fiber.Ctx) (err error) {
	return
}

func (u *AuthHandler) Login(c *fiber.Ctx) error {
	input := new(domain.Login)

	// parse body
	if err := c.BodyParser(input); err != nil {
		return errUtil.Result(c, errUtil.ERROR, fiber.StatusUnprocessableEntity, "review your input body", err.Error())
	}

	// validator.v9
	if err := isRequestValid(input); err != nil {
		return errUtil.Result(c, errUtil.ERROR, fiber.StatusUnprocessableEntity, "review your input body", err.Error())
	}

	// using input email
	findUser, err := u.UserUsecase.GetByEmail(c.Context(), input.Email)
	if err != nil {
		return errUtil.Result(c, errUtil.ERROR, fiber.StatusInternalServerError, "user not found", err.Error())
	}

	// hashing
	if !CheckPasswordHash(input.Password, findUser.Password) {
		return errUtil.Result(c, errUtil.ERROR, fiber.StatusInternalServerError, "user password hash error", nil)
	}

	zap.L().Info("find user", zap.Int64("id", findUser.ID))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = findUser.Username
	claims["user_id"] = findUser.ID
	//claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(u.secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return errUtil.Result(c, errUtil.SUCCESS, fiber.StatusOK, "found", t)

	//return c.JSON(fiber.Map{"status": errUtil.SUCCESS, "code": fiber.StatusOK, "message": "found", "data": t})
}
