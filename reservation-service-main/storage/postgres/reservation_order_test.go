package postgres

import (
	"log"
	"testing"

	pb "github.com/dilshodforever/reservation-service/genprotos"
	"github.com/stretchr/testify/assert"
)

func TestCreateReservationOrder(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	ReservationOrder := &pb.ReservationOrder{
		Id:            "UUID",
		ReservationId: "UUID",
		MenuItemId:    "UUID",
		Quantity:      "19",
	}

	result, err := stg.ReservationOrder().CreateReservationOrder(ReservationOrder)

	assert.NoError(t, err)
	assert.NotNil(t, result)

}

func TestGetByIdReservationOrder(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	user, err := stg.ReservationOrder().GetByIdReservationOrder(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetAllReservationOrder(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	ReservationOrders, err := stg.ReservationOrder().GetAllReservationOrder(&pb.ReservationOrder{})
	assert.NoError(t, err)
	assert.NotNil(t, ReservationOrders)
}

func TestUpdateReservationOrder(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	ReservationOrder := &pb.ReservationOrder{
		Id:            "UUID",
		ReservationId: "UUID",
		MenuItemId:    "UUID",
		Quantity:      "19",
	}

	result, err := stg.ReservationOrder().UpdateReservationOrder(ReservationOrder)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestDeleteReservationOrder(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}

	var Id pb.ById

	Id.Id = "b409ff53-ff2b-4033-84b4-4ce555081647"

	result, err := stg.ReservationOrder().DeleteReservationOrder(&Id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
