package entity_test

import (
	"testing"
	"yoyoj-go-rag/internal/user/domain/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := entity.NewUser("testname", "123456@qq.com", "123456")

	assert.NoError(t, err)
	assert.Equal(t, "testname", user.Username)
	assert.Equal(t, "123456@qq.com", user.Email)
	assert.Equal(t, "123456", user.Password)
}
