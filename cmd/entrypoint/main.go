package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/ppzxc/golang-boilerplate-in-my-case/proc"
	"github.com/ppzxc/golang-boilerplate-in-my-case/util/config/logger"
	"github.com/ppzxc/golang-boilerplate-in-my-case/util/config/yml"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

const (
	packageName    = "test"
	processName    = "testd"
	processVersion = "0.0.1"
)

var (
	configFilePath  = flag.String("c", "", "absolute path of the configuration file")
	logLevel        = flag.String("l", "INFO", "log level")
	logFilePath     = flag.String("p", "", "log file path")
	dsn             = flag.String("d", "", "dsn")
	goMaxProcessNum = flag.Int("g", 0, "GOMAXPROCS number")
	//timeZone        = flag.String("tz", "Asia/Seoul", "set time zone")
)

// @title golang fiber boilerplate in my case
// @version 0.0.1
// @description This is boilerplate using fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email cjh8487@naver.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// timezone
	//_, _ = time.LoadLocation(*timeZone)
	//_ = time.Now().In(loc)

	//go func() {
	//	app := fiber.New()
	//	app.Get("/docs/*", swagger.Handler) // default
	//	//app.Get("/docs/*", swagger.New(swagger.Config{ // custom
	//	//	URL: "http://example.com/doc.json",
	//	//	DeepLinking: false,
	//	//}))
	//	app.Listen(":8080")
	//}()

	// flags
	flag.Parse()

	// MaxProcess
	if *goMaxProcessNum == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	} else {
		runtime.GOMAXPROCS(*goMaxProcessNum)
	}

	// validation, config file path
	if len(*configFilePath) <= 0 {
		fmt.Println("-c config file path is invalid")
		os.Exit(-1)
	}
	fmt.Printf("CONFIG FILE PATH : %+v\r\n", *configFilePath)

	// validation, log level
	if len(*logLevel) <= 0 {
		fmt.Println("-l loglevel is not set")
		os.Exit(-1)
	}
	fmt.Printf("LOG LEVEL : %+v\r\n", *logLevel)

	// validation, sentry dsn
	if len(*dsn) <= 0 {
		fmt.Println("-d sentry dsn is not set")
		os.Exit(-1)
	}
	fmt.Printf("SENTRY DSN : %+v\r\n", *dsn)
	fmt.Printf("LOG FILE PATH : %+v\r\n", *logFilePath)
	fmt.Printf("USE LOG FILE : %+v\r\n", len(*logFilePath) > 0)

	// init, logger
	if err := logger.Init(*logLevel, *logFilePath, len(*logFilePath) > 0, *dsn); err != nil {
		zap.L().Error("logger init error occurred", zap.Error(err))
		os.Exit(-1)
	}

	// init, config file
	config, err := yml.Read(*configFilePath, *logLevel, *logFilePath, len(*logFilePath) < 1, *dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// shutdown context
	c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	shutdownCtx, cancel := context.WithCancel(context.Background())

	go func() {
		osCall := <-c
		zap.L().Warn("OS Signal Received",
			zap.String("signal", osCall.String()))
		cancel()
	}()

	zap.L().Info("process start",
		zap.String("package", packageName),
		zap.String("process", processName),
		zap.String("version", processVersion))

	// call main
	if err := proc.Main(shutdownCtx, config); err != nil {
		zap.L().Error("main process is shutdown",
			zap.Error(err))
	}
}
