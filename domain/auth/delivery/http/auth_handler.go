package http

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/ppzxc/golang-boilerplate-in-my-case/domain"
	"github.com/ppzxc/golang-boilerplate-in-my-case/middleware"
	errUtil "github.com/ppzxc/golang-boilerplate-in-my-case/util/err"
	validator "gopkg.in/go-playground/validator.v9"
	"time"
)

type AuthHandler struct {
	//AuthUsecase domain.AuthUsecase
	UserUsecase domain.UserUsecase
}

//func NewAuthHandler(r fiber.Router, au domain.AuthUsecase, uu domain.UserUsecase) {

func NewAuthHandler(r fiber.Router, uu domain.UserUsecase, secret string) {
	handler := &AuthHandler{
		//AuthUsecase: au,
		UserUsecase: uu,
	}

	user := r.Group("/auth")
	user.Post("/login", handler.Login)
	user.Post("/verify", middleware.Protected(secret), handler.Verify)
}

func isRequestValid(m *domain.Login) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *AuthHandler) Verify(c *fiber.Ctx) (err error) {
	return
}

func (u *AuthHandler) Login(c *fiber.Ctx) (err error) {
	var inputLogin domain.Login
	err = c.BodyParser(&inputLogin)
	if err != nil {
		return errUtil.Result(c, errUtil.ERROR, fiber.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&inputLogin); !ok {
		return errUtil.Result(c, errUtil.ERROR, fiber.StatusUnprocessableEntity, err.Error())
	}

	dataLogin, err := u.UserUsecase.GetByEmail(c.Context(), inputLogin.Email)
	if err != nil {
		return errUtil.Result(c, errUtil.ERROR, fiber.StatusInternalServerError, err.Error())
	}

	if inputLogin.Email != dataLogin.Email || inputLogin.Password != dataLogin.Password {
		return errUtil.Result(c, errUtil.ERROR, fiber.StatusUnauthorized, "unauthorized")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = inputLogin.Email
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": errUtil.SUCCESS, "code": fiber.StatusOK, "message": "found", "data": t})
}
