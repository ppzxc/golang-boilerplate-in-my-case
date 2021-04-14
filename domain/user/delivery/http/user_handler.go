package http

import (
	"github.com/ppzxc/golang-boilerplate-in-my-case/domain"
	"github.com/ppzxc/golang-boilerplate-in-my-case/middleware"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)
import "github.com/gofiber/fiber/v2"

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(r fiber.Router, uu domain.UserUsecase, secret string) {
	handler := &UserHandler{
		UserUsecase: uu,
	}

	user := r.Group("/user")
	user.Post("/", handler.CreateUser)

	user.Get("/:id", middleware.Protected(secret), handler.GetByID)
	user.Delete("/:id", middleware.Protected(secret), handler.DeleteUser)
	user.Post("/:id", middleware.Protected(secret), handler.UpdateUser)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//func validToken(t *jwt.Token, id string) bool {
//	n, err := strconv.Atoi(id)
//	if err != nil {
//		return false
//	}
//
//	claims := t.Claims.(jwt.MapClaims)
//	uid := int(claims["user_id"].(float64))
//
//	if uid != n {
//		return false
//	}
//
//	return true
//}

//func validUser(id string, p string) bool {
//	db := database.DB
//	var user model.User
//	db.First(&user, id)
//	if user.Username == "" {
//		return false
//	}
//	if !CheckPasswordHash(p, user.Password) {
//		return false
//	}
//	return true
//}

func (u *UserHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// validation
	if len(id) <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "code": fiber.StatusBadRequest, "message": "/user/:id parameter is invalid", "data": nil})
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "code": fiber.StatusBadRequest, "message": "/user/:id parameter is invalid", "data": nil})
	}

	// get by id
	byID, err := u.UserUsecase.GetByID(c.Context(), int64(intId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "code": fiber.StatusBadRequest, "message": "/user/:id parameter is invalid", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "code": fiber.StatusOK, "message": "Product found", "data": byID})
}

// CreateUser new user
func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	//type NewUser struct {
	//	Username string `json:"username"`
	//	Email    string `json:"email"`
	//}
	//
	//db := database.DB
	//user := new(model.User)
	//if err := c.BodyParser(user); err != nil {
	//	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	//
	//}
	//
	//hash, err := hashPassword(user.Password)
	//if err != nil {
	//	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	//
	//}
	//
	//user.Password = hash
	//if err := db.Create(&user).Error; err != nil {
	//	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	//}
	//
	//newUser := NewUser{
	//	Email:    user.Email,
	//	Username: user.Username,
	//}

	//return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
	return c.JSON(fiber.Map{"status": "success", "code": fiber.StatusOK, "message": "Created user", "data": "d"})
}

// UpdateUser update user
func (u *UserHandler) UpdateUser(c *fiber.Ctx) error {
	//type UpdateUserInput struct {
	//	Names string `json:"names"`
	//}
	//var uui UpdateUserInput
	//if err := c.BodyParser(&uui); err != nil {
	//	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	//}
	//id := c.Params("id")
	//token := c.Locals("user").(*jwt.Token)
	//
	//if !validToken(token, id) {
	//	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	//}
	//
	//db := database.DB
	//var user model.User
	//
	//db.First(&user, id)
	//user.Names = uui.Names
	//db.Save(&user)

	//return c.JSON(fiber.Map{"status": "success", "message": "User successfully updated", "data": user})
	return c.JSON(fiber.Map{"status": "success", "code": fiber.StatusOK, "message": "User successfully updated", "data": "d"})
}

// DeleteUser delete user
func (u *UserHandler) DeleteUser(c *fiber.Ctx) error {
	//type PasswordInput struct {
	//	Password string `json:"password"`
	//}
	//var pi PasswordInput
	//if err := c.BodyParser(&pi); err != nil {
	//	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	//}
	//id := c.Params("id")
	//token := c.Locals("user").(*jwt.Token)
	//
	//if !validToken(token, id) {
	//	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	//
	//}
	//
	//if !validUser(id, pi.Password) {
	//	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})
	//
	//}
	//
	//db := database.DB
	//var user model.User
	//
	//db.First(&user, id)
	//
	//db.Delete(&user)
	return c.JSON(fiber.Map{"status": "success", "code": fiber.StatusOK, "message": "User successfully deleted", "data": nil})
}
