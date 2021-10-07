package users_test

import (
	"context"
	"fmt"
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

var DB users2.Repository
func TestMain(m *testing.M) {
	db := initDB()
	DB = users.NewRepoMySQL(db)
	m.Run()
	defer tearDown(db)
}

func TestRepoUsers_StoreNewUsers(t *testing.T) {
	t.Run("success add", func(t *testing.T) {
		res := &users2.Domain{
			Email: "avtarakhrisna1@gmail.com",
			Photo: "testing.com",
		}
		res, err := DB.StoreNewUsers(context.Background(), res)

		assert.NotNil(t, res)
		assert.NoError(t, err)
	})

	t.Run("duplicate", func(t *testing.T) {
		res := &users2.Domain{
			Email: "avtarakhrisna1@gmail.com",
			Photo: "testing.com",
		}
		_, err := DB.StoreNewUsers(context.Background(), res)

		assert.NotNil(t, err)
	})
}

func TestRepoUsers_GetByUserID(t *testing.T) {
	t.Run("no record", func(t *testing.T) {
		uuidUserID, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")
		res, err := DB.GetByUserID(context.Background(), uuidUserID);

		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("find record", func(t *testing.T) {
		data := &users2.Domain{
			Name: "avtarakhrisna12@gmail.com",
			Photo: "testing.com",
		}
		res, _ := DB.StoreNewUsers(context.Background(), data)
		userUUID := res.UserID
		result, err := DB.GetByUserID(context.Background(), userUUID)

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

func TestRepoUsers_GetByEmail(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		data := &users2.Domain{
			Email: "avtarakhrisna9090@gmail.com",
			Photo: "testing.com",
		}
		res, _ := DB.StoreNewUsers(context.Background(), data)
		res, err := DB.GetByEmail(context.Background(), "avtarakhrisna9090@gmail.com")

		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("not found", func(t *testing.T) {
		res, err := DB.GetByEmail(context.Background(), "kamil@gmail.com")

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
}

func TestRepoUsers_GetByEmailAndPassword(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		res, err := DB.GetByEmailAndPassword(context.Background(), "kamil@gmail.com")

		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("found", func(t *testing.T) {
		data := &users2.Domain{
			Email: "avtarakhrisna909ss0@gmail.com",
			Photo: "testing.com",
		}
		res, _ := DB.StoreNewUsers(context.Background(), data)
		res, err := DB.GetByEmailAndPassword(context.Background(), "avtarakhrisna909ss0@gmail.com")

		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}

func TestRepoUsers_UpdateStatus(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		data := &users2.Domain{
			Email: "avtarakhrisna909ss0dwd@gmail.com",
			Photo: "testing.com",
		}
		res, _ := DB.StoreNewUsers(context.Background(), data)
		err := DB.UpdateStatus(context.Background(), res.UserID)

		assert.NoError(t, err)
	})
}