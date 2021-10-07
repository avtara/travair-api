package reservations_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/avtara/travair-api/businesses/reservations"
	"github.com/avtara/travair-api/businesses/units"
	users2 "github.com/avtara/travair-api/businesses/users"
	_reservRepo "github.com/avtara/travair-api/repository/databases/reservations"
	_unitsRepo "github.com/avtara/travair-api/repository/databases/units"
	"github.com/avtara/travair-api/repository/databases/users"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

func initDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		"localhost", "avtara", "avtara112", "testing", "5432")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err.Error())
	}
	dbMigrate(db)
	return db
}

func dbMigrate(db *gorm.DB) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.AutoMigrate(&users.Users{}, &_unitsRepo.Units{}, &_unitsRepo.Photos{}, &_reservRepo.Reservation{})
}

func tearDown(db *gorm.DB) {
	db.Migrator().DropTable(&users.Users{}, &_unitsRepo.Units{}, &_unitsRepo.Photos{}, &_reservRepo.Reservation{}, &_unitsRepo.Address{})
}

var DB reservations.Repository
var UsersDB users2.Repository
var UnitDB units.Repository
func TestMain(m *testing.M) {
	db := initDB()
	DB = _reservRepo.NewRepoMySQL(db)
	UsersDB = users.NewRepoMySQL(db)
	UnitDB = _unitsRepo.NewRepoMySQL(db)
	m.Run()
	defer tearDown(db)
}

func TestRepoUnit_Store(t *testing.T) {
	t.Run("success add", func(t *testing.T) {
		fromDate, _ := time.Parse("2006-01-02", "2006-01-02")
		endDate, _ := time.Parse("2006-01-02", "2006-01-03")
		usr := &users2.Domain{
			Email: "dwkokd@gmail.com",
			Photo: "testing.com",
		}
		userRes, _ := UsersDB.StoreNewUsers(context.Background(), usr)
		unt := &units.Domain{
			Name: "awdok",
			OwnerID: userRes.ID,
			Address: units.Address{
				Street: "awd",
			},
		}
		untStr, _ := UnitDB.Store(context.Background(),unt)
		resv := &reservations.Domain{
			CustomerID: userRes.ID,
			CheckInDate: fromDate,
			CheckOutDate: endDate,
			UnitID:untStr.ID,
		}
		result, err := DB.Store(context.Background(), resv)

		assert.NotNil(t, result)
		assert.NoError(t, err)
	})

	t.Run("fail store", func(t *testing.T) {
		fromDate, _ := time.Parse("2006-01-02", "2006-01-02")
		endDate, _ := time.Parse("2006-01-02", "2006-01-03")
		resv := &reservations.Domain{
			CheckInDate: fromDate,
			CheckOutDate: endDate,
		}
		result, err := DB.Store(context.Background(), resv)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})
}

func TestRepoUnit_GetByDate(t *testing.T) {
	t.Run("success get", func(t *testing.T) {
		fromDate, _ := time.Parse("2006-01-02", "2006-01-02")
		endDate, _ := time.Parse("2006-01-02", "2006-01-03")
		resv := &reservations.Domain{
			CheckInDate: fromDate,
			CheckOutDate: endDate,
		}
		result := DB.GetByDate(context.Background(), resv)

		assert.NotNil(t, result)
	})

	t.Run("not found", func(t *testing.T) {
		fromDate, _ := time.Parse("2006-01-02", "2006-01-02")
		endDate, _ := time.Parse("2006-01-09", "2006-01-03")
		resv := &reservations.Domain{
			CheckInDate: fromDate,
			CheckOutDate: endDate,
		}
		result := DB.GetByDate(context.Background(), resv)

		assert.Equal(t, errors.New("record not found"), result)
	})
}
