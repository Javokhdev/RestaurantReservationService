package postgres

import (
	"log"
	"testing"

	pb "github.com/dilshodforever/restaurant-auth/genprotos"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	user := &pb.User{
		Id:       "b409ff53-ff2b-4033-84b4-4ce555081647",
		UserName: "Mukhammed_777",
		Password: "password123",
		Email:    "unique_test_user@example.com",
	}
	result, err := stg.User().CreateUser(user)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestGetByIdUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	user, err := stg.User().GetByIdUser(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetAllUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	users, err := stg.User().GetAllUser(&pb.User{})
	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestUpdateUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	user := &pb.User{
		Id:       "user_id_here",
		UserName: "updated_user",
		Email:    "updated_user@example.com",
	}
	result, err := stg.User().UpdateUser(user)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestDeleteUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	result, err := stg.User().DeleteUser(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestLoginUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	user, err := stg.User().LoginUser(&pb.User{})

	assert.NoError(t, err)
	assert.NotNil(t, user)
}
