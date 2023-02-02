package usecases_test

import (
	"api/src/user"
	"api/src/user/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUserHappyCase(t *testing.T) {
	assert := assert.New(t)
	userMemRepo := user.UserInMemoryRepo{UserList: make([]domain.User, 0), IsErr: false}

}
