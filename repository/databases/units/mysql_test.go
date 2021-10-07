package units_test

import (
	"context"
	"fmt"
	"github.com/avtara/travair-api/businesses/units"
	users2 "github.com/avtara/travair-api/businesses/users"
	_reservRepo "github.com/avtara/travair-api/repository/databases/reservations"
	_unitsRepo "github.com/avtara/travair-api/repository/databases/units"
	"github.com/avtara/travair-api/repository/databases/users"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
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

var UsersDB users2.Repository
var UnitDB units.Repository
var unitUUID uuid.UUID
var unitID uint
var userID uint
var untDomain *units.Domain

func TestMain(m *testing.M) {
	db := initDB()
	UsersDB = users.NewRepoMySQL(db)
	UnitDB = _unitsRepo.NewRepoMySQL(db)
	userData := &users2.Domain{
		Email: "dwkokd@gmail.com",
		Photo: "testing.com",
		Role: "tenant",
		Status: 1,
	}
	userResult, _ := UsersDB.StoreNewUsers(context.Background(), userData)

	userID = userResult.ID
	untDomain = &units.Domain{
		Name:    "Hotel Mas Joko",
		OwnerID: userResult.ID,
		Address: units.Address{
			Street: "Jalan Gunung Sumbing",
			City: "Purwokerto",
			Country: "Indonesia",
			Latitude: -7.4245,
			Longitude: 109.2302,
			PostalCode: "53202",
			State: "Central Java",
		},
	}
	untStr, _ := UnitDB.Store(context.Background(), untDomain)

	unitUUID = untStr.UnitID
	unitID, _ = UnitDB.GetIDByUnitID(context.Background(), unitUUID)
	m.Run()
	defer tearDown(db)
}

func TestRepoUnit_Store(t *testing.T) {
	t.Run("fail store", func(t *testing.T) {
		unt := &units.Domain{
			Name: "awdok",
		}
		untStr, err := UnitDB.Store(context.Background(), unt)

		assert.Nil(t, untStr)
		assert.Error(t, err)
	})
}

func TestRepoUnit_GetIDByUnitID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		untStr, err := UnitDB.GetIDByUnitID(context.Background(), unitUUID)

		assert.NotNil(t, untStr)
		assert.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		uuidFake, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")
		untStr, err := UnitDB.GetIDByUnitID(context.Background(), uuidFake)

		assert.Equal(t,uint(0), untStr)
		assert.Error(t, err)
	})
}

func TestRepoUnit_UpdatePathByUnitID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := UnitDB.UpdatePathByUnitID(context.Background(), unitUUID, "localhost:8080/blabla")

		assert.Nil(t, err)
		assert.NoError(t, err)
	})
}

func TestRepoUnit_AddPhotoUnit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := UnitDB.AddPhotoUnit(context.Background(), unitID, "localhost:8080/blabla")

		assert.Nil(t, err)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		err := UnitDB.AddPhotoUnit(context.Background(), 0, "localhost:8080/blabla")

		assert.NotNil(t, err)
		assert.Error(t, err)
	})
}

func TestRepoUnit_SelectAllPhotosByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		_, err := UnitDB.SelectAllPhotosByID(context.Background(), unitID)
		assert.Nil(t, err)
		assert.NoError(t, err)
	})
}

func TestRepoUnit_GetByUnitID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		res, err := UnitDB.GetByUnitID(context.Background(), unitUUID)
		assert.Nil(t, err)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}

func TestRepoUnit_SelectAddressByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		res, err := UnitDB.SelectAddressByID(context.Background(), unitID)
		assert.Nil(t, err)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}

func TestRepoUnit_GetUnitsByGeo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		res, err := UnitDB.GetUnitsByGeo(context.Background(), -7.4245, 109.2302)
		assert.Nil(t, err)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}
