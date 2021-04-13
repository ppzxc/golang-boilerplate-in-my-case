package proc

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ppzxc/golang-boilerplate-in-my-case/user/delivery/http"
	"github.com/ppzxc/golang-boilerplate-in-my-case/user/repository/mariadb"
	"github.com/ppzxc/golang-boilerplate-in-my-case/user/usecase"
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

	// fiber init
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	// context timeout
	timeoutContext := time.Duration(config.Http.Context.Timeout) * time.Second

	// user init
	ur := mariadb.NewMariadbUserRepository(dbConn)
	uu := usecase.NewUserUsecase(ur, timeoutContext)
	http.NewUserHandler(app, uu)

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
