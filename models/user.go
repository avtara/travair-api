package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

//Role product type
type Role string

//const available value for enum
const (
	Tenant Role = "tenant"
	Guest  Role = "guest"
)

//Value validate enum when set to database
func (t Role) Value() (driver.Value, error) {
	switch t {
	case Tenant, Guest: //valid case
		return string(t), nil
	}
	return nil, errors.New("Invalid product type value") //else is invalid
}

//Scan validate enum on read from data base
func (t *Role) Scan(value interface{}) error {
	var pt Role
	if value == nil {
		*t = ""
		return nil
	}
	st, ok := value.(string) // if we declare db type as ENUM gorm will scan value as []uint8
	if !ok {
		return errors.New("invalid data for product type")
	}
	pt = Role(st) //convert type from string to ProductType
	switch pt {
	case Tenant, Guest: //valid case
		*t = pt
		return nil
	}
	return fmt.Errorf("invalid product type value :%s", st) //else is invalid
}

//User represents users table in database
type User struct {
	gorm.Model
	ID          uint64    `gorm:"primary_key:auto_increment"`
	UserID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string    `gorm:"type:varchar(255)"`
	Username    string    `gorm:"uniqueIndex;type:varchar(255)"`
	Email       string    `gorm:"uniqueIndex;type:varchar(255)"`
	Password    string    `gorm:"->;<-;not null" `
	Phone       string    `gorm:"uniqueIndex;type:varchar(255)" `
	Photo       string    `gorm:"type:varchar(255)"`
	Role        Role      `sql:"type:user_access"`
	Status      int       `gorm:"default:0;size:10"`
	DateOfBirth datatypes.Date
}
