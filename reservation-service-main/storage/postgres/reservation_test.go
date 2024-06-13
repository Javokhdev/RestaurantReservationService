package postgres

import (
	"log"
	"testing"

	pb "github.com/dilshodforever/reservation-service/genprotos"
	"github.com/stretchr/testify/assert"
)

func TestCreateReservation(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	Reservation := &pb.Reservation{
		Id:           "b409ff53-ff2b-4033-84b4-4ce555081647",
		ReservationTime: "2024-06-22",
		Status: "Ok",
	}

	result, err := stg.Reservation().CreateReservation(Reservation)

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

	user, err := stg.Reservation().GetByIdReservation(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetAllReservation(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	Reservations, err := stg.Reservation().GetAllReservation(&pb.Reservation{})
	assert.NoError(t, err)
	assert.NotNil(t, Reservations)
}

func TestUpdateReservation(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	Reservation := &pb.Reservation{
		Id:           "b409ff53-ff2b-4033-84b4-4ce555081647",
		//RestaurantId: "b409ff53-ff2b-4033-84b4-4ce555081647",
		ReservationTime: "2024-06-22",
		Status: "Ok",
	}

	result, err := stg.Reservation().UpdateReservation(Reservation)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestDeleteReservation(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	result, err := stg.Reservation().DeleteReservation(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

