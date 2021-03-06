package main

import (
	"github.com/avtara/travair-api/app/config"
	_middleware "github.com/avtara/travair-api/app/middleware"
	"github.com/avtara/travair-api/app/routes"
	_reservationsService "github.com/avtara/travair-api/businesses/reservations"
	_unitsService "github.com/avtara/travair-api/businesses/units"
	_usersService "github.com/avtara/travair-api/businesses/users"
	_reservationsController "github.com/avtara/travair-api/controllers/reservations"
	_unitsController "github.com/avtara/travair-api/controllers/units"
	_usersController "github.com/avtara/travair-api/controllers/users"
	"github.com/avtara/travair-api/helpers"
	_reservationsRepo "github.com/avtara/travair-api/repository/databases/reservations"
	_unitsRepo "github.com/avtara/travair-api/repository/databases/units"
	_usersRepo "github.com/avtara/travair-api/repository/databases/users"
	_queueRepo "github.com/avtara/travair-api/repository/queue"
	"github.com/avtara/travair-api/repository/thirdparties/ipapi"
	"github.com/avtara/travair-api/repository/uploads/local"
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
	)
	port := os.Getenv("PORT")
	timeJWT, _ := strconv.Atoi(os.Getenv("JWT_TOKEN_AGE"))
	secretToken := os.Getenv("SECRET_TOKEN_KEY")
	baseUrl := os.Getenv("BASE_URL")
	timeoutDur, _ := strconv.Atoi(os.Getenv("TIMEOUT_IN_MS"))
	timeoutContext := time.Duration(timeoutDur) * time.Millisecond
	configJWT := _middleware.ConfigJWT{
		SecretJWT:       secretToken,
		ExpiresDuration: timeJWT,
	}

	e := echo.New()
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(_middleware.LoggerConfig()))
	e.Use(middleware.Static("/assets"))

	queueRepo := _queueRepo.NewRepoAMPQ(ampq)
	uploadRepo := local.NewUploadRepository("assets", baseUrl+":"+port)
	ipapiRepo := ipapi.NewIpAPI()

	userRepo := _usersRepo.NewRepoMySQL(db)
	userService := _usersService.NewUserService(userRepo, timeoutContext, queueRepo, &configJWT)
	userCtrl := _usersController.NewUserController(userService, ampq)

	unitRepo := _unitsRepo.NewRepoMySQL(db)
	unitService := _unitsService.NewUnitService(unitRepo, userService, timeoutContext, uploadRepo, ipapiRepo)
	unitCtrl := _unitsController.NewUnitController(unitService)

	reservationsRepo := _reservationsRepo.NewRepoMySQL(db)
	reservationsService := _reservationsService.NewReservationService(reservationsRepo, userService, timeoutContext, unitService)
	reservationsCtrl := _reservationsController.NewReservationController(reservationsService)

	routesInit := routes.ControllerList{
		JWTMiddleware:  configJWT.Init(),
		UserController: *userCtrl,
		UnitController: *unitCtrl,
		ReservationController : *reservationsCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":" + port))
}
