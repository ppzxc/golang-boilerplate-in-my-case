package proc

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ppzxc/golang-boilerplate-in-my-case/domain/auth/delivery/http"
	http2 "github.com/ppzxc/golang-boilerplate-in-my-case/domain/user/delivery/http"
	mariadb2 "github.com/ppzxc/golang-boilerplate-in-my-case/domain/user/repository/mariadb"
	usecase2 "github.com/ppzxc/golang-boilerplate-in-my-case/domain/user/usecase"
	"github.com/ppzxc/golang-boilerplate-in-my-case/util/config/yml"
	custom "github.com/ppzxc/golang-boilerplate-in-my-case/util/err"
	"go.uber.org/zap"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Main(ctx context.Context, config *yml.Config) error {
	// example data base connection
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DataBase.Username,
		config.DataBase.Password,
		config.DataBase.Host,
		config.DataBase.Port,
		config.DataBase.Instance)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Seoul")
	dsn := fmt.Sprintf("%s?%s", connectionString, val.Encode())
	dbConn, err := sql.Open(config.DataBase.Type, dsn)

	if err != nil {
		zap.L().Fatal("data base connection open error occurred", zap.Error(err))
	}
	err = dbConn.Ping()
	if err != nil {
		zap.L().Fatal("data base ping error occurred", zap.Error(err))
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			zap.L().Fatal("data base connection close error occurred", zap.Error(err))
		}
	}()

	fiberConfig := fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return ctx.Status(code).JSON(fiber.Map{"status": "error", "code": code, "message": err.Error(), "data": nil})
		},
	}

	// fiber init
	app := fiber.New(fiberConfig)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	api := app.Group("/api", logger.New())
	v1 := api.Group("/v1")

	// context timeout
	timeoutContext := time.Duration(config.Http.Context.Timeout) * time.Second

	// user init
	ur := mariadb2.NewMariadbUserRepository(dbConn)
	uu := usecase2.NewUserUsecase(ur, timeoutContext)
	http2.NewUserHandler(v1, uu, config.Http.Jwt.Secret)

	http.NewAuthHandler(v1, uu, config.Http.Jwt.Secret)

	// fiber run
	go func() {
		if err := app.Listen(config.Http.Addr); err != nil {
			zap.L().Panic("fiber panic!", zap.Error(err))
		}
	}()

	// fiber signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	zap.L().Info("start main process")
	for {
		select {
		case <-ctx.Done():
			if err := app.Shutdown(); err != nil {
				zap.L().Error("fiber shutdown error", zap.Error(err))
			}
			zap.L().Warn("main process is shutdown...")
			return custom.MainProcessContextTerminated
		case <-c:
			if err := app.Shutdown(); err != nil {
				zap.L().Error("fiber shutdown error", zap.Error(err))
				return err
			}
			zap.L().Warn("main process is shutdown...")
			return nil
			//default:
			//	zap.L().Info("hi",
			//		zap.String("config", fmt.Sprintf("%+v", config)))
			//	time.Sleep(1 * time.Second)
		}
	}
}
