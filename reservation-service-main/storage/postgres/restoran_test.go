package postgres

import (
	"log"
	"testing"

	pb "github.com/dilshodforever/reservation-service/genprotos"
	"github.com/stretchr/testify/assert"
)

func TestCreateRestoran(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	Restoran := &pb.Restoran{
		Id:          "UUID",
		Name:        "JizBiz",
		Address:     "Toshkent",
		PhoneNumber: "9059449494",
		Description: "Very good",
	}

	result, err := stg.Restoran().CreateRestoran(Restoran)

	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestGetByIdRestoran(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	user, err := stg.Restoran().GetByIdRestoran(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetAllRestoran(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	Restorans, err := stg.Restoran().GetAllRestoran(&pb.Restoran{})
	assert.NoError(t, err)
	assert.NotNil(t, Restorans)
}

func TestUpdateRestoran(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	Restoran := &pb.Restoran{
		Id:          "UUID",
		Name:        "JizBiz",
		Address:     "Toshkent",
		PhoneNumber: "9059449494",
		Description: "Very good",
	}

	result, err := stg.Restoran().UpdateRestoran(Restoran)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestDeleteRestoran(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	result, err := stg.Restoran().DeleteRestoran(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
