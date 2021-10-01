package helpers

import (
	"github.com/labstack/gommon/random"
)

func RandomToken(n int) (string) {
	data := random.String(uint8(n))
	return data
}
