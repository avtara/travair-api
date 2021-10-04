package helpers_test

import (
	"github.com/avtara/travair-api/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_Registration(t *testing.T) {
	t.Run("Success generate random token", func(t *testing.T) {
		res := helpers.RandomToken(2)
		assert.Equal(t, len(res), 2)
	})

	t.Run("Encrypt success", func(t *testing.T) {
		res, err := helpers.HashPassword("using")
		t.Skip(res,err)
	})


}

