package main

import (
	"github.com/avtara/travair-api/app/config"
	_middleware "github.com/avtara/travair-api/app/middleware"
	"github.com/avtara/travair-api/app/routes"
	"github.com/avtara/travair-api/businesses/users"
	_usersController "github.com/avtara/travair-api/controllers/users"
	"github.com/avtara/travair-api/helpers"
	"github.com/avtara/travair-api/repository/databases/cache"

	//"github.com/avtara/travair-api/repository/databases/cache"
	_usersRepo "github.com/avtara/travair-api/repository/databases/users"
	"github.com/avtara/travair-api/repository/queue"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	var (
		db   = config.SetupDatabaseConnection()
		ampq = config.SetupAMPQConnection()
		rdb  = config.SetupRedisConnection()
	)
	timeoutDur, _ := strconv.Atoi(os.Getenv("TIMEOUT_CONTEXT"))
	timeoutContext := time.Duration(timeoutDur) * time.Second

	e := echo.New()
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(_middleware.LoggerConfig()))

	queueRepo := queue.NewRepoAMPQ(ampq)

	cacheRepo := cache.NewRepoCache(rdb)

	userRepo := _usersRepo.NewRepoMySQL(db)
	userService := users.NewUserService(userRepo, timeoutContext, queueRepo, cacheRepo)
	userCtrl := _usersController.NewUserController(userService, ampq)

	routesInit := routes.ControllerList{
		UserController: *userCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8080"))
}
