package database

import (
	"errors"
	"fmt"
	"github.com/avtara/travair-api/config"
	"github.com/avtara/travair-api/models"
	"github.com/avtara/travair-api/utils"
	"github.com/jackc/pgconn"
	"strings"
)

func InsertUser(user *models.User) (*models.User, error) {
	var err error
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	if results :=  config.DB.Create(&user); results.Error != nil {
		if pgError := results.Error.(*pgconn.PgError); errors.Is(results.Error, pgError) {
			switch pgError.Code {
			case "23505":
				columnName, _ := utils.GetDetailColumnSQL(pgError.Detail)
				return nil, errors.New(fmt.Sprintf("This %s is already in use!", strings.Title(columnName)))
			}

			return nil, errors.New(pgError.Message)
		}
	}
	return user, nil
}